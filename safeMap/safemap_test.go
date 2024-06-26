package safemap

import (
	"math/rand"
	"testing"
)

// TestNewAndSet, Yeni ve Ayarla işlevlerini test eder
func TestNewAndSet(t *testing.T) {
	sm := New[int, string]() // yeni bir safemap oluştur
	sm.Set(1, "one")         // safemap'a bir anahtar ve bir deger ekler
	sm.Set(2, "two")         // ve tekrar ekler

	if sm.Len() != 2 { // safemap'in uzunluğunu test eder
		t.Errorf("Expected length 2, got %d", sm.Len()) // hata mesajını verir
	}
}

// TestGet, Okuma işlevini test eder
func TestGet(t *testing.T) {
	sm := New[int, string]() // yeni bir safemap oluştur
	sm.Set(1, "one")         // safemap'a bir anahtar ve bir deger ekler

	val, ok := sm.Get(1)     // safemap'ten bir anahtar alır
	if !ok || val != "one" { // anahtarın varlığını ve degerini test eder
		t.Errorf("Expected 'one', got %s", val) // hata mesajını verir
	}
	_, ok = sm.Get(2) // anahtarın varlığını test eder
	if ok {           // anahtarın varlığını test eder
		t.Errorf("Expected false, got true") // hata mesajını verir
	}
}

// TestDelete, Silme işlevini test eder
func TestDelete(t *testing.T) {
	sm := New[int, string]() // yeni bir safemap oluştur
	sm.Set(1, "one")         // safemap'a bir anahtar ve bir deger ekler
	sm.Delete(1)             // anahtarı siler
	_, ok := sm.Get(1)       // anahtarın varlığını test eder
	if ok {                  // anahtarın varlığını test eder
		t.Errorf("Expected key 1 to be deleted") // hata mesajını verir
	}

	sm.Delete(2) // just make sure this doesn't panic
}

// TestLen, Uzunluk bilgisini test eder
func TestLen(t *testing.T) {
	sm := New[int, string]() // yeni bir safemap oluştur
	sm.Set(1, "one")         // safemap'a bir anahtar ve bir deger ekler
	sm.Set(2, "two")         // ve tekrar ekler
	sm.Set(3, "three")       // ve tekrar ekler
	sm.Delete(2)             // anahtarı siler

	if sm.Len() != 2 { // safemap'in uzunluğunu test eder
		t.Errorf("Expected length 2, got %d", sm.Len()) // hata mesajını verir
	}
}

// TestForEach, ForEach işlevini test eder
func TestForEach(t *testing.T) {
	sm := New[int, string]() // yeni bir safemap oluştur
	sm.Set(1, "one")         // safemap'a bir anahtar ve bir deger ekler
	sm.Set(2, "two")         // ve tekrar ekler

	keys := make([]int, 0)             // anahtarları tutacak dizi oluştur
	sm.ForEach(func(k int, v string) { // ForEach işlevini test eder
		keys = append(keys, k) // anahtarı diziye ekler
	})

	if len(keys) != 2 { // dizi uzunlugunu test eder
		t.Errorf("Expected 2 keys, got %d", len(keys)) // hata mesajını verir
	}

	// Check if keys 1 and 2 are present
	if !contains(keys, 1) || !contains(keys, 2) { // anahtarın varlığını test eder
		t.Errorf("Expected keys 1 and 2, got %v", keys) // hata mesajını verir
	}
}

// Helper function to check if a slice contains a specific element.
func contains(slice []int, element int) bool {
	for _, a := range slice { // slice'deki her bir elemanı döngüde kontrol eder
		if a == element { // eşleşme bulunduğunda
			return true // true döndürür
		}
	}
	return false // yoksa false döndürür
}

// Benchmark Get
func BenchmarkGetConcurrent(b *testing.B) {
	ds := New[uint64, uint64]() // Yeni bir güvenli (safe) map oluştur
	// Test verileriyle veri deposunu önceden doldurmak gerekiyorsa
	for i := 0; i < 100000; i++ {
		ds.Set(uint64(i), uint64(i)) // Güvenli mape bir anahtar ve bir değer ekle
	}
	b.SetParallelism(100)                // Paralel olarak çalışacak gorutin sayısını ayarla
	b.ResetTimer()                       // Benchmark zamanlayıcısını sıfırla
	b.RunParallel(func(pb *testing.PB) { // Benchmark'ı paralel olarak çalıştır
		for pb.Next() { // Her iterasyon için
			r := rand.Uint64() % 100000 // Rastgele bir anahtar üret
			_, _ = ds.Get(r)            // Değeri al
		}
	})
}

// Compiler'ı aldat
var silly uint64

// Benchmark Set
func BenchmarkSetConcurrent(b *testing.B) {
	ds := New[uint64, uint64]() // Yeni bir güvenli (safe) map oluştur
	b.SetParallelism(100)       // Paralel olarak çalışacak gorutin sayısını ayarla
	b.ResetTimer()              // Benchmark zamanlayıcısını sıfırla

	b.RunParallel(func(pb *testing.PB) { // Benchmark'ı paralel olarak çalıştır
		for pb.Next() { // Her iterasyon için
			r := rand.Uint64() % 100000 // Rastgele bir anahtar üret
			silly, _ = ds.Get(r)        // Değeri al
			ds.Set(r, r)                // Değeri ayarla
			silly, _ = ds.Get(r)        // Değeri tekrar al
		}
	})
}
