## 코드 분석

### package main
모든 코드는 package로 시작된다.  
Go의 모든 코드는 어떤 package안에 포함되어야한다.
#### main은 프로그램의 시작점을 포함하는 패키지라는 의미를 가지고 있다. 또한 단 하나이다.
시작점이란 무엇일까? 프로그램이 시작되면 메모리에 로드된다.
CPU가 이것을 읽어서 실행시킨다. 여기서 문제는 "어디가 시작일까?"이다.
그래서 첫번째 줄에 시작점을 적어 둔다.(프로그램 스타팅 포인트 => main함수)
package main은 이 프로그램의 스타팅 포인트를 포함된다.

### import "fmt"
자기 자신을 가져올 수 없다.

### func main() {
함수를 선언하는 명령어
main 함수는 시작점을 의미하는 특별한 명령어이다.
main 함수가 즉 프로그램의 스타팅 포인트이다.