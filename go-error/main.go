// package main

import "fmt"

// PasswordError struct 선언
type PasswordError struct {
	Len int
	RequireLen int
}

// PasswordError 객체에 속한, Error 메서드 생성
// 이로 인해 PasswordError 객체는 Error라는 메서드를 가지고 있게 되기 때문에
// golang에서 제공하는 기본적인 error 타입으로서 사용할 수 있게 된다.
func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

// 회원가입하는 함수 - 입력한 패스워드의 길이가 8보다 짧으면 PasswordError 객체를 return한다.
// 이때 PasswordError 객체는 Error()메서드를 가지고 있기 때문에 error 타입으로 리턴이 가능하다.
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil // nil도 error로 들어갈 수 있기 때문에 이 함수에서 return 값으로 사용 가능
}

func main() {
	var id string
	var pwd string
	fmt.Scan(&id, &pwd)
	err := RegisterAccount(id, pwd)
	if err != nil {
		// 아래에서는 err 객체를 PasswordError 객체로 변환한다. 이 때 err객체는 error 타입이고
		// PasswordError 객체는 Error() 메서드를 포함하고 있기 때문에 error타입으로 사용이 가능하다.
		// 그래서 둘 사이의 변환이 가능하다.
		if errInfo, ok := err.(PasswordError); ok { // ok에는 변환의 성공 여부가 담긴다.
			// 여기서 이상한 것이, errInfo에는 PasswordError 타입 객체가 담기는데
			// errInfo만 써놓아도 errInfo.Error()의 결과값이 출력되는 것이다.
			// 이 이유는 추후에 알아봐야 할 것 같다.
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo.Error(), errInfo.Len, errInfo.RequireLen)
		} 
	} else {
		fmt.Println("회원가입 되었습니다.")
	}
}