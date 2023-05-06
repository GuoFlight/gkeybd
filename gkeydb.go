package gkeybd

import (
	"errors"
	"fmt"
	"github.com/micmonay/keybd_event"
)

// 按下字符串
func TypeStr(str string) error {

	if len(str) == 0 {
		return nil
	}
	// 创建键盘事件模拟器
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	// 轮循按下按键
	for _, char := range str {
		if vks, ok := KbMap[string(char)]; ok {
			//是否按下shift
			if _, ok := KbShouldShift[string(char)]; ok {
				kb.HasSHIFTR(true)
			}

			kb.SetKeys(vks...)
			err = kb.Launching()
			kb.HasSHIFTR(false) //重置Shift
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("The key '%c' is undefined", char))
		}
	}
	return nil
}
