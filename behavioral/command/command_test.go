package command

import "testing"

func TestCommand(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		tv := &Tv{}

		onCommand := &OnCommand{
			device: tv,
		}

		offCommand := &OffCommand{
			device: tv,
		}

		onButton := &Button{
			command: onCommand,
		}
		onButton.press()

		offButton := &Button{
			command: offCommand,
		}
		offButton.press()

		///
		radio := &Radio{}

		onCommand2 := &OnCommand{
			device: radio,
		}

		offCommand2 := &OffCommand{
			device: radio,
		}

		onButton2 := &Button{
			command: onCommand2,
		}
		onButton2.press()

		offButton2 := &Button{
			command: offCommand2,
		}
		offButton2.press()
	})
}
