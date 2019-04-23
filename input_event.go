package term

var (
	inputEvtCh     = make(chan Event, 64)
	inputPrepareCh = make(chan int, 1)
)

func inputPrepare() {
	select {
	case inputPrepareCh <- 1:
	default:
	}
}

func inputEvt(evtCh <-chan Event) {
	var es []Event
	for {
		select {
		case e := <-evtCh:
			if len(es) > 0 {
				es = append(es, e) // TODO deal withs OOM
			}

			select {
			case inputEvtCh <- e:
			default:
				es = append(es, e)
			}

		case <-inputPrepareCh:
			if len(es) == 0 {
				break
			}
			e := es[0]
			select {
			case inputEvtCh <- e:
				es = es[1:]
			default:
			}
		case <-inputQuit:
			return
		}
	}
}
