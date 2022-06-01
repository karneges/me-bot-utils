package utils

func ConcurrencyMap[T interface{}, E interface{}](arr []T, cb func(arg T) (E, error), concurrency int) ([]E, error) {
	var store []E
	inputCh := make(chan T)
	outputCh := make(chan E)
	errorCh := make(chan error)
	if len(arr) == 0 {
		return store, nil
	}
	go func() {
		for idx, nft := range arr {
			inputCh <- nft
			if idx == len(arr)-1 {
				close(inputCh)
			}
		}
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			for arrEl := range inputCh {
				cbResult, err := cb(arrEl)
				if err != nil {
					errorCh <- err
				}
				outputCh <- cbResult
			}
		}()

	}

	for {
		select {
		case newValue := <-outputCh:
			store = append(store, newValue)
			if len(store) == len(arr) {
				return store, nil
			}
		case err := <-errorCh:
			return store, err
		}
	}
}

func WithRetry[R interface{}](executable func() (R, error), retryCount int) (R, error) {
	if retryCount == 0 {
		return executable()
	}
	res, err := executable()
	if err != nil {
		return WithRetry(executable, retryCount-1)
	}
	return res, nil
}

func Pipe2[P1 interface{}, T1 interface{}, T2 *interface{}](f1 func(P1) (T1, error), f2 func(T1) (T2, error)) func(P1) (T2, error) {
	return func(p1 P1) (T2, error) {
		result1, err := f1(p1)
		if err != nil {
			return nil, err
		}
		return f2(result1)
	}
}
func Pipe3[P1 interface{}, T1 interface{}, T2 interface{}, T3 *interface{}](
	f1 func(P1) (T1, error),
	f2 func(T1) (T2, error),
	f3 func(T2) (T3, error),
) func(P1) (T3, error) {
	return func(p1 P1) (T3, error) {
		result1, err := f1(p1)
		if err != nil {
			return nil, err
		}
		result2, err := f2(result1)
		if err != nil {
			return nil, err
		}
		return f3(result2)
	}
}
func Pipe4[P1 any, T1 any, T2 any, T3 any, T4 any](
	f1 func(P1) (T1, error),
	f2 func(T1) (T2, error),
	f3 func(T2) (T3, error),
	f4 func(T3) (T4, error),
) func(P1) (T4, error) {
	var t T4

	return func(p1 P1) (T4, error) {
		result1, err := f1(p1)
		if err != nil {
			return t, err
		}
		result2, err := f2(result1)
		if err != nil {
			return t, err
		}
		result3, err := f3(result2)
		if err != nil {
			return t, err
		}
		return f4(result3)
	}
}
func Pipe5[P1 any, T1 any, T2 any, T3 any, T4 any, T5 any](
	f1 func(P1) (T1, error),
	f2 func(T1) (T2, error),
	f3 func(T2) (T3, error),
	f4 func(T3) (T4, error),
	f5 func(T4) (T5, error),
) func(P1) (T5, error) {
	var t T5
	return func(p1 P1) (T5, error) {
		result1, err := f1(p1)
		if err != nil {
			return t, err
		}
		result2, err := f2(result1)
		if err != nil {
			return t, err
		}
		result3, err := f3(result2)
		if err != nil {
			return t, err
		}
		result4, err := f4(result3)
		if err != nil {
			return t, err
		}
		return f5(result4)
	}
}
func Pipe6[P1 interface{}, T1 interface{}, T2 interface{}, T3 interface{}, T4 interface{}, T5 interface{}, T6 *interface{}](
	f1 func(P1) (T1, error),
	f2 func(T1) (T2, error),
	f3 func(T2) (T3, error),
	f4 func(T3) (T4, error),
	f5 func(T4) (T5, error),
	f6 func(T5) (T6, error),
) func(P1) (T6, error) {
	return func(p1 P1) (T6, error) {
		result1, err := f1(p1)
		if err != nil {
			return nil, err
		}
		result2, err := f2(result1)
		if err != nil {
			return nil, err
		}
		result3, err := f3(result2)
		if err != nil {
			return nil, err
		}
		result4, err := f4(result3)
		if err != nil {
			return nil, err
		}
		result5, err := f5(result4)
		if err != nil {
			return nil, err
		}
		return f6(result5)
	}
}
