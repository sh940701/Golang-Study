package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	// 한 단어씩 읽는 bufio패키지의 Scanner 생성, io.Reader인터페이스를 인수로 받기 때문에
	// string 타입을 io.Reader로 만들어주기 위해 strings.NewReader()함수를 사용했다.
	// 문자열을 한 줄씩 혹은 한 단어씩 끊어 읽고자 할 때 주로 사용되는 구문이다.
	scanner := bufio.NewScanner(strings.NewReader(str))

	// Scanner객체의 Split()메서드를 호출해, 어떻게 끊어 읽을지 알려준다.
	// bufil.ScanWord -> 한 단어씩, bufio.ScanLines -> 한 줄씩
	scanner.Split(bufio.ScanWords)

	// 여기서 pos는, 에러를 반환할 때, 입력받은 문자열 어디에서 문제가 생겼는지를 알려주기 위한 변수
	pos := 0

	a, n, err := readNextInt(scanner)
	if err != nil {
		// readNextInt()에서 발생한 에러를 다시 감싸서 pos 정보를 에러에 추가했다.
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1

	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	// scanner 객체의 Scan() 메서드를 사용해 첫번째 단어를 읽어온다. 이 때 if문에 쓰였지만
	// scanner.Scan()을 실행한 순간, scanner에는 읽어온 단어가 저장된다. 결과값은 성공여부 boolean
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Fail to scan")
	}
	
	// 위에서 읽어온 한 단어의 데이터를 Text 메서드를 사용해 변수에 저장해준다.
	word := scanner.Text()

	// strconv.Atoi()메서드는 문자열을 int 타입으로 변경한다 <->Itoa
	// 변경시  숫자가 아닌 문자가 섞여있는 경우에는 NumberError타입의 에러를 반환한다.
	number, err := strconv.Atoi(word)

	// 위에서 발생한 에러를 감쌌다. 이로 인해 string -> number 과정에서의 에러임을 알 수 있다.
	// 이 때 fmt.Errorf()의 서식문자 %w를 사용하면 err를 감싸서 새로운 에러를 반환하게 된다.
	// 따라서 다른 정보와 strconv.Atoi() 함수에서 발생한 에러까지 에러 하나로 반환할 수 있다.
	if err != nil {
		return 0, 0, fmt.Errorf("Filed to convert word to int, word:%s err:%w", word, err)
	}

	// 에러가 없을 시 읽어온 숫자와 그 길이, nil을 반환해준다.
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		// errors.As() 함수는, err 안에 감싸진 에러 중 두 번째 인수의 타입인 *strconv.NumError
		// 로 변환될 수 있는 에러가 있다면 변환하여 값을 넣고 true를 반환하는 함수이다.
		// 우리가 감싼 에러가 strconv.Atoi()함수에서 발생한 에러이기 때문에 이 에러는
		// *strconv.NumError타입이라서, errors.As() 함수는 true를 반환하고 numError에
		// *strconv.NumError 타입값을 넣어준다. 따라서 감싸진 에러를 검사하고 각 에러 타입별로
		// 다른 처리를 해줄 수가 있다.
		if errors.As(err, &numError) {
			fmt.Println("NumberError", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
