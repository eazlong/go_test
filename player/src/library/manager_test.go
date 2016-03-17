package library

import "testing"

func TestOper( t *testing.T ) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error( "mm is nil" )
	}
	
	m0 := &MusicEntry{ "one day" }
	mm.Add( m0 );
	if mm.Len() != 1 {
		t.Error( "mm len is not ", mm.Len )
	}
	
	m := mm.Find( m0.Name )
	if m == nil {
		t.Error( "Find fail" )
	}
	
	m, err := mm.GetAt(0)
	if m == nil {
		t.Error( err )
	}
	
	m = mm.Remove( 0 )
	if m == nil || mm.Len() != 0 {
		t.Error( "mm Remove failed", err )
	}
}