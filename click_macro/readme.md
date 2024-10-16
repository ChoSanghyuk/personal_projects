# go 매크로 작성

robotgo 라이브러리를 사용해서, 예약 도움 도구 만드는 프로젝트

사용 라이브러리
gohook
robotgo 

사전 설치
gcc (window에서는 tdm gcc 설치 https://jmeubank.github.io/tdm-gcc/)

사용팁
마우스 이동 시에는 중간중간 time.Sleep 필요 (0.3초 권장)
이벤트 등록 시, 키보드 이벤트는 `hook.AddEvent`, 마우스 이벤트는 `hook.AddMouse`
