## Test Code
파일명이 _test.go로 끝나야한다.
testing 패키지를 임포트해야 합니다.
테스트 코드는 func TestXxxx(t *testing.T) 형태이어야 합니다.

### 일부만 테스트하기 
go test -run { 테스트명 }

### 테스트를 돕는 외부 패키지
go get github.com/stretchr/testify

### 테스트 주도 개발(TDD)
테스트가 중요한데 테스트 케이스가 절대적으로 부족하거나 형식적인 경우가 많았다. => 테스트가 부족해지는 경우가 발생
이유로는 코드 작성이 끝나고 나서 예상 가능한 테스트 케이스만 만들어 테스트하는 경우가 많았기 때문에

테스트 주도 개발이란? 

테스트부터 먼저 작성 -> 테스트 실패 -> 코드 작성 -> 테스트 성공 -> 개선
이것을 완성까지 반복

장점은 정말 많다 그래서 단점에 대해서 소개하겠다.

단점
1. 모듈간의 의존성이 높은 경우에 테스트 케이스를 만들기 힘들다.

    해결방법
        - 의존성을 끊는 방법
        - Mock Up을 만들어서 테스트 한다.

2. 동시성 테스트에 취약하다. (멀티 스레드에서 테스트가 힘들다.)
3. 진정한 TDD가 아닌 형식적인 테스트로 전락할 수 있다. (사람이 하는 일이라 아무래도 대충하면 대충만들어진다.)
4. 지속적인 모니터링과 관리가 필요하다. (교육이 필요하다.)

## 벤치마크
성능 검사시 사용

다음과 같은 표현 규약
1. 파일명이 _test.go로 끝나야한다.
2. testing 패키지를 임포트해야한다.
3. 밴치마크 코드는 func BenchmarkXxxx(b *testing.B)형태이여야한다.
4. go test -bench { 테스트명 }의 명령어로 동작한다.