package impl

type nbChan[T any] chan T

func (ch nbChan[T]) send(v T) (dropped bool) {
	for {
		select {
		case ch <- v:
			return dropped
		default:
			select {
			case <-ch:
				dropped = true
			default:
			}
		}
	}
}
