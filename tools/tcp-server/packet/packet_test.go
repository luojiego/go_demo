package packet

import "testing"

func TestSubmitAck_Decode(t *testing.T) {
	b := []byte{'0', '0', '0', '0', '0', '0', '0', '1', '1'}
	s := &SubmitAck{}
	if err := s.Decode(b); err != nil {
		t.Error("decode err: ", err)
	}

	if s.ID != "00000001" {
		t.Errorf("id: %s not eq 00000001", s.ID)
	}

	if s.Result != '1' {
		t.Errorf("result: %c not eq 1", s.Result)
	}
}
