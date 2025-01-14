package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/formancehq/ledger/pkg/api/apierrors"
	"github.com/formancehq/ledger/pkg/core"
	"github.com/formancehq/ledger/pkg/ledger/command"
	"github.com/formancehq/ledger/pkg/storage/ledgerstore"
	sharedapi "github.com/formancehq/stack/libs/go-libs/api"
	"github.com/formancehq/stack/libs/go-libs/errorsutil"
	"github.com/formancehq/stack/libs/go-libs/metadata"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func CountTransactions(w http.ResponseWriter, r *http.Request) {
	l := LedgerFromContext(r.Context())

	var startTimeParsed, endTimeParsed core.Time
	var err error
	if r.URL.Query().Get(QueryKeyStartTime) != "" {
		startTimeParsed, err = core.ParseTime(r.URL.Query().Get(QueryKeyStartTime))
		if err != nil {
			apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation, ErrInvalidStartTime))
			return
		}
	}

	if r.URL.Query().Get(QueryKeyEndTime) != "" {
		endTimeParsed, err = core.ParseTime(r.URL.Query().Get(QueryKeyEndTime))
		if err != nil {
			apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation, ErrInvalidEndTime))
			return
		}
	}

	txQuery := ledgerstore.NewTransactionsQuery().
		WithReferenceFilter(r.URL.Query().Get("reference")).
		WithAccountFilter(r.URL.Query().Get("account")).
		WithSourceFilter(r.URL.Query().Get("source")).
		WithDestinationFilter(r.URL.Query().Get("destination")).
		WithStartTimeFilter(startTimeParsed).
		WithEndTimeFilter(endTimeParsed).
		WithMetadataFilter(sharedapi.GetQueryMap(r.URL.Query(), "metadata"))

	count, err := l.CountTransactions(r.Context(), txQuery)
	if err != nil {
		apierrors.ResponseError(w, r, err)
		return
	}

	w.Header().Set("Count", fmt.Sprint(count))
	sharedapi.NoContent(w)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	l := LedgerFromContext(r.Context())

	txQuery := ledgerstore.NewTransactionsQuery()

	if r.URL.Query().Get(QueryKeyCursor) != "" {
		if r.URL.Query().Get("after") != "" ||
			r.URL.Query().Get("reference") != "" ||
			r.URL.Query().Get("account") != "" ||
			r.URL.Query().Get("source") != "" ||
			r.URL.Query().Get("destination") != "" ||
			r.URL.Query().Get(QueryKeyStartTime) != "" ||
			r.URL.Query().Get(QueryKeyEndTime) != "" ||
			r.URL.Query().Get(QueryKeyPageSize) != "" {
			apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
				errors.Errorf("no other query params can be set with '%s'", QueryKeyCursor)))
			return
		}

		err := ledgerstore.UnmarshalCursor(r.URL.Query().Get(QueryKeyCursor), &txQuery)
		if err != nil {
			apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
				errors.Errorf("invalid '%s' query param", QueryKeyCursor)))
			return
		}
	} else {
		var (
			err             error
			afterTxIDParsed uint64
		)
		if r.URL.Query().Get("after") != "" {
			afterTxIDParsed, err = strconv.ParseUint(r.URL.Query().Get("after"), 10, 64)
			if err != nil {
				apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
					errors.New("invalid 'after' query param")))
				return
			}
		}

		var startTimeParsed, endTimeParsed core.Time
		if r.URL.Query().Get(QueryKeyStartTime) != "" {
			startTimeParsed, err = core.ParseTime(r.URL.Query().Get(QueryKeyStartTime))
			if err != nil {
				apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation, ErrInvalidStartTime))
				return
			}
		}

		if r.URL.Query().Get(QueryKeyEndTime) != "" {
			endTimeParsed, err = core.ParseTime(r.URL.Query().Get(QueryKeyEndTime))
			if err != nil {
				apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation, ErrInvalidEndTime))
				return
			}
		}

		pageSize, err := getPageSize(r)
		if err != nil {
			apierrors.ResponseError(w, r, err)
			return
		}

		txQuery = txQuery.
			WithAfterTxID(afterTxIDParsed).
			WithReferenceFilter(r.URL.Query().Get("reference")).
			WithAccountFilter(r.URL.Query().Get("account")).
			WithSourceFilter(r.URL.Query().Get("source")).
			WithDestinationFilter(r.URL.Query().Get("destination")).
			WithStartTimeFilter(startTimeParsed).
			WithEndTimeFilter(endTimeParsed).
			WithMetadataFilter(sharedapi.GetQueryMap(r.URL.Query(), "metadata")).
			WithPageSize(pageSize)
	}

	cursor, err := l.GetTransactions(r.Context(), txQuery)
	if err != nil {
		apierrors.ResponseError(w, r, err)
		return
	}

	sharedapi.RenderCursor(w, *cursor)
}

type Script struct {
	core.Script
	Vars map[string]any `json:"vars"`
}

func (s Script) ToCore() core.Script {
	s.Script.Vars = map[string]string{}
	for k, v := range s.Vars {
		switch v := v.(type) {
		case string:
			s.Script.Vars[k] = v
		case map[string]any:
			s.Script.Vars[k] = fmt.Sprintf("%s %v", v["asset"], v["amount"])
		default:
			s.Script.Vars[k] = fmt.Sprint(v)
		}
	}
	return s.Script
}

type PostTransactionRequest struct {
	Postings  core.Postings     `json:"postings"`
	Script    Script            `json:"script"`
	Timestamp core.Time         `json:"timestamp"`
	Reference string            `json:"reference"`
	Metadata  metadata.Metadata `json:"metadata" swaggertype:"object"`
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	l := LedgerFromContext(r.Context())

	payload := PostTransactionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		apierrors.ResponseError(w, r,
			errorsutil.NewError(command.ErrValidation,
				errors.New("invalid transaction format")))
		return
	}

	if len(payload.Postings) > 0 && payload.Script.Plain != "" ||
		len(payload.Postings) == 0 && payload.Script.Plain == "" {
		apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
			errors.New("invalid payload: should contain either postings or script")))
		return
	} else if len(payload.Postings) > 0 {
		if i, err := payload.Postings.Validate(); err != nil {
			apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation, errors.Wrap(err,
				fmt.Sprintf("invalid posting %d", i))))
			return
		}
		txData := core.TransactionData{
			Postings:  payload.Postings,
			Timestamp: payload.Timestamp,
			Reference: payload.Reference,
			Metadata:  payload.Metadata,
		}

		res, err := l.CreateTransaction(r.Context(), getCommandParameters(r), core.TxToScriptData(txData))
		if err != nil {
			apierrors.ResponseError(w, r, err)
			return
		}

		sharedapi.Ok(w, res)
		return
	}

	script := core.RunScript{
		Script:    payload.Script.ToCore(),
		Timestamp: payload.Timestamp,
		Reference: payload.Reference,
		Metadata:  payload.Metadata,
	}

	res, err := l.CreateTransaction(r.Context(), getCommandParameters(r), script)
	if err != nil {
		apierrors.ResponseError(w, r, err)
		return
	}

	sharedapi.Ok(w, res)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	l := LedgerFromContext(r.Context())

	txId, err := strconv.ParseUint(chi.URLParam(r, "txid"), 10, 64)
	if err != nil {
		apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
			errors.New("invalid transaction ID")))
		return
	}

	tx, err := l.GetTransaction(r.Context(), txId)
	if err != nil {
		apierrors.ResponseError(w, r, err)
		return
	}

	sharedapi.Ok(w, tx)
}

func RevertTransaction(w http.ResponseWriter, r *http.Request) {
	l := LedgerFromContext(r.Context())

	txId, err := strconv.ParseUint(chi.URLParam(r, "txid"), 10, 64)
	if err != nil {
		apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
			errors.New("invalid transaction ID")))
		return
	}

	tx, err := l.RevertTransaction(r.Context(), getCommandParameters(r), txId)
	if err != nil {
		apierrors.ResponseError(w, r, err)
		return
	}

	sharedapi.Created(w, tx)
}

func PostTransactionMetadata(w http.ResponseWriter, r *http.Request) {
	l := LedgerFromContext(r.Context())

	var m metadata.Metadata
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
			errors.New("invalid metadata format")))
		return
	}

	txId, err := strconv.ParseUint(chi.URLParam(r, "txid"), 10, 64)
	if err != nil {
		apierrors.ResponseError(w, r, errorsutil.NewError(command.ErrValidation,
			errors.New("invalid transaction ID")))
		return
	}

	if err := l.SaveMeta(r.Context(), getCommandParameters(r), core.MetaTargetTypeTransaction, txId, m); err != nil {
		apierrors.ResponseError(w, r, err)
		return
	}

	sharedapi.NoContent(w)
}
