/* checkip - A remote IP check API written in Go
 *
 * Author: Thomas Guyot-Sionnest <Thomas@Guyot-Sionnest.net>
 *
 * Copyright 2018 Thomas Guyot-Sionnest
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package checkip

import (
	"fmt"
	"strings"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Disable caching on proxies, etc.
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	w.Header().Set("Content-Type", "text/plain")
	/* XFF is only used in dev mode (we don't use load balancers), and
	 * even then it's hard to differenciate between what's sent by the
	 * user/proxy and intermediates... Best would be to take first
	 * element and thus trust XFF's added by proxies! Commented out
	 * for that reason
	*/
	xff, xffPresent := r.Header["X-Forwarded-For"]
	if xffPresent {
		// Just use the first header (closest one)
		fmt.Fprint(w, strings.Split(xff[0], ", ")[0], "\n")
	} else {
		fmt.Fprint(w, r.RemoteAddr, "\n")
	}
}

