# search

### Available Operations

* [search](#search) - Search
* [searchgetServerInfo](#searchgetserverinfo) - Get server info

## search

ElasticSearch query engine

### Example Usage

```java
package hello.world;

import com.formance.formance_sdk.SDK;
import com.formance.formance_sdk.models.operations.SearchResponse;
import com.formance.formance_sdk.models.shared.Query;
import com.formance.formance_sdk.models.shared.Security;

public class Application {
    public static void main(String[] args) {
        try {
            SDK sdk = SDK.builder()
                .setSecurity(new Security("mollitia") {{
                    authorization = "Bearer YOUR_ACCESS_TOKEN_HERE";
                }})
                .build();

            com.formance.formance_sdk.models.shared.Query req = new Query() {{
                after = new String[]{{
                    add("users:002"),
                    add("users:002"),
                }};
                cursor = "YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol=";
                ledgers = new String[]{{
                    add("quickstart"),
                    add("quickstart"),
                }};
                pageSize = 221262L;
                policy = "OR";
                raw = new java.util.HashMap<String, Object>() {{
                    put("odit", "nemo");
                    put("quasi", "iure");
                    put("doloribus", "debitis");
                    put("eius", "maxime");
                }};
                sort = "txid:asc";
                target = "deleniti";
                terms = new String[]{{
                    add("destination=central_bank1"),
                    add("destination=central_bank1"),
                    add("destination=central_bank1"),
                }};
            }};            

            SearchResponse res = sdk.search.search(req);

            if (res.response != null) {
                // handle response
            }
        } catch (Exception e) {
            // handle exception
        }
    }
}
```

## searchgetServerInfo

Get server info

### Example Usage

```java
package hello.world;

import com.formance.formance_sdk.SDK;
import com.formance.formance_sdk.models.operations.SearchgetServerInfoResponse;
import com.formance.formance_sdk.models.shared.Security;

public class Application {
    public static void main(String[] args) {
        try {
            SDK sdk = SDK.builder()
                .setSecurity(new Security("in") {{
                    authorization = "Bearer YOUR_ACCESS_TOKEN_HERE";
                }})
                .build();

            SearchgetServerInfoResponse res = sdk.search.searchgetServerInfo();

            if (res.serverInfo != null) {
                // handle response
            }
        } catch (Exception e) {
            // handle exception
        }
    }
}
```
