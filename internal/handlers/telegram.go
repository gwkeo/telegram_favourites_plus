package handlers

//import (
//	"github.com/gwkeo/telegram_favourites_plus/internal/models"
//	"sync"
//)
//
//type Type int
//
//const (
//	Text Type = iota
//	Animation
//)
//
//func HandleTelegramResponse(results []models.Result) {
//	wg := &sync.WaitGroup{}
//	for _, r := range results {
//		wg.Add(1)
//		go func() {
//			resType := TypeOfResult(r, wg)
//
//		}()
//	}
//
//	wg.Wait()
//}
//
//func TypeOfResult(r models.Result, wg *sync.WaitGroup) Type {
//	defer wg.Done()
//	if r.Message.Text != "" {
//		return Text
//	} else {
//		return Animation
//	}
//}
