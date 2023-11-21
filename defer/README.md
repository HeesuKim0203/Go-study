## Defer 문

Todo : 활용 방법 찾아보기

ex9.1
함수가 종료 될때까지 실행을 연기한다. 하지만 인자는 즉시 평가된다.

ex9.2
연기된 함수 호출들은 스택에 쌓인다. 한 함수가 종료될 때 연기된 함수들은 후입선출 순서로 수행된다. 

참고 자료 : [https://go.dev/blog/defer-panic-and-recover](https://go.dev/blog/defer-panic-and-recover)