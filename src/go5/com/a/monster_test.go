package main
import "testing"

func TestStore(t *testing.T){
	var mon Monster = Monster{Name:"小三",Age:23,Skill:"喷火",}
	mon.Store()
	t.Logf("good")
}

func TestReStore(t *testing.T){
	var mon Monster
	mon.ReStore()
	t.Logf("姓名=%v,年龄=%v,技能=%v",mon.Name,mon.Age,mon.Skill)
	if mon.Name == "小三"{
		t.Logf("good")
	}else{
		t.Fatalf(mon.Name)
	}
}
