# Whiteblock logger package

화이트블록의 기본 로거입니다

## Usage
### 1. 로거 인스턴스 사용
```go
package main

import (

    "context"
    "github.com/whiteblockco/go-pkg/logger"
    "github.com/whiteblockco/go-pkg/logger/stackdriver"
    "github.com/whiteblockco/go-pkg/logger/stdout"
    "gopkg.in/errgo.v2/errors"
 
)

func main() {
    var lgr logger.Logger
    var err error
    lgr, err = stackdriver.New(context.Background(), "", "", "")
    if err != nil {
        lgr = stdout.New()
    }

    // Start logging
    lgr.Info("blabla")
    
    lgr.ErrorWarning(errors.New("this is new error"))
}
```

### 2. 전역 로거 사용
```go
package main

import (

    "context"
    "github.com/whiteblockco/go-pkg/logger"
    "github.com/whiteblockco/go-pkg/logger/stackdriver"
    "github.com/whiteblockco/go-pkg/logger/stdout"
    "gopkg.in/errgo.v2/errors"
 
)

func main() {
    var lgr logger.Logger
    var err error
    lgr, err = stackdriver.New(context.Background(), "", "", "")
    if err != nil {
        lgr = stdout.New()
    }

    // 전역 로거 설정
    logger.SetLogger(lgr)
}

func someFunction() {
    // Start logging
    logger.Info("blabla")
        
    logger.ErrorWarning(errors.New("this is new error"))
}
```

### 로그 종류
#### Info
단순한 정보를 기록하기 위해 사용합니다. string 값을 기록합니다.


#### Warning
경고 수준의 로그를 생성합니다. string 값을 기록합니다.


#### Critical
심각 수준의 로그를 생성합니다. string 값을 기록합니다.


#### Data
데이터를 덤프하기 위해 사용합니다. 다음과 같은 경우에 사용할 수 있습니다.
- 400 (BadRequest)일 경우 RequestBody 덤프
- 에러가 발생한 곳에서 메모리 내용 덤프
- 메타데이터 기록 (데이터 분석을 위함)


#### LogTrace
Trace정보가 포함된 로그를 생성합니다.  
stackdriver의 추적(지연시간)을 위한 API이며 사용하지 않아도 무방합니다.  
**trace정보를 같이 제공하여 합니다.**


#### LogHTTP
http Request (net/http의 Request)를 덤프하여 기록합니다.  
trace정보가 존재하면 자동으로 추출하여 추적도 같이 기록합니다.


#### ErrorWarning
경고 수준의 error인터페이스를 리포팅 합니다.   
로그가 아닌 stackdriver의 ErrorReporting (오류 보고)에 stacktrace와 함께 기록됩니다.  
**제공하는 error가 stacktrace 정보를 포함하고 있어야 stacktrace까지 기록됩니다**  
tip:  errors.New()로 error 정보를 생성하면 자동으로 stacktrace정보를 포함합니다.  


#### ErrorCritical 
심각 수준의 error인터페이스를 리포팅 합니다.   
로그가 아닌 stackdriver의 ErrorReporting (오류 보고)에 stacktrace와 함께 기록됩니다.



## Report Issue (Story, Proposal, Bug, etc)
깃허브 이슈 탭에 이슈를 만들어주세요.
제안은 언제나 환영입니다.

제안 시 예시 코드도 만들어 주면 더욱 땡큐
