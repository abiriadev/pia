<h1 align="center">Pia</h1>
<p align="center">노벨피아 비공식 API</p>

> [!CAUTION]
> 본 API는 ㈜메타크래프트사에서 보장하는 공식 기능이 아니며, 본 API를 사용함으로 인하여 발생하는 모든 법적인 책임은 본 API의 사용자가 부담합니다. 또한, 이 API를 저작권 침해, 불법 배포 등에 절대로 악용하지 말아 주세요.

# 사용

```sh
$ go get -u github.com/abiriadev/pia@latest
```

# 예시

```go
package main

import (
	"fmt"

	"github.com/abiriadev/pia"
)

func main() {
	client := pia.NewPiaClient("")

	content, err := client.Novel(1720152).Content()
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
```

# 지원 API

-   [ ] 작품 목록 조회
    -   [ ] 텍스트 검색
    -   [ ] 태그 검색
-   [ ] 작품 정보 조회
    -   [ ] 표지
    -   [ ] 태그
    -   [ ] 추천수
-   [ ] 작품 회차 조회
-   [x] 작품 본문 조회
    -   [ ] 유저 댓글 조회
-   [ ] 작가 조회

# 라이선스

[![GitHub](https://img.shields.io/github/license/abiriadev/pia?color=7632ff&style=for-the-badge)](./LICENSE)
