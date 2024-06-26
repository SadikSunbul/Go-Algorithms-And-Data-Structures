package safemap

import "sync"

/*
"K" türü, derleyici tarafından karşılaştırma operatörleri (örneğin <, ==, >, vb.) ile karşılaştırılabilir olmalıdır yani (int,flout,vs.).
V any: Bu ifade, "V" türünün herhangi bir tür olabileceğini belirtir. Yani "V" türü, herhangi bir veri türü olabilir.
*/

type SafeMap[K comparable, V any] struct { //Generic type kullandık
	mu sync.RWMutex
	/*
		sync.RWMutex Go programlama dilinde bir senkronizasyon mekanizmasıdır. "RW" kısaltması "read-write" anlamına
		gelir, yani bu mutex (kilit) hem okuma (reading) hem de yazma (writing) işlemlerini senkronize etmek için kullanılır.

		Eşzamanlı Erişimi Kontrol Etmek: Bir RWMutex nesnesi, belirli bir kaynağa eşzamanlı erişimi kontrol etmek için
		kullanılır. Eğer bir iş parçacığı (goroutine) bu kaynağa yazma işlemi yapacaksa, diğer iş
		parçacıklarının bu kaynağa herhangi bir erişim yapmasına izin verilmez (ne okuma ne de yazma).

		Okuma (Reading) İşlemlerini Paralel Yapmak: RWMutex özellikle birden fazla iş parçacığının aynı anda belirli bir
		kaynağı okumasına izin vermek için kullanılır. Birden fazla iş parçacığı aynı anda okuma işlemi yapabilirken, yazma
		işlemi yapacak iş parçacığı bu esnada bekler.
	*/
	data map[K]V
}

// Yeni bir SafeMap nesnesi oluşturur ve döndürer.
func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{ // safemap nesnesi oluşturuldu
		data: make(map[K]V), // map oluşturuldu
	}
}

// Yeni bir anahtar ekleme işlemini gerçekleştir
func (s *SafeMap[K, V]) Set(k K, v V) {
	s.mu.Lock()         // mutex kilitlendi
	defer s.mu.Unlock() // mutex kilitlendi
	s.data[k] = v       // anahtar eklendi
}

// Okuma işlemini gerçekleştir
func (s *SafeMap[K, V]) Get(k K) (V, bool) {
	s.mu.RLock()         // mutex kilitlendi
	defer s.mu.RUnlock() // mutex kilitlendi
	val, ok := s.data[k] // anahtar okundu
	return val, ok
}

// Anahtar silme işlemini gerçekleştir
func (s *SafeMap[K, V]) Delete(k K) {
	s.mu.Lock()         // mutex kilitlendi
	defer s.mu.Unlock() // mutex kilitlendi
	delete(s.data, k)   // anahtar silindi
}

// Map'in boyutunu gerçekleştir
func (s *SafeMap[K, V]) Len() int {
	s.mu.RLock()         // mutex kilitlendi
	defer s.mu.RUnlock() // mutex kilitlendi
	return len(s.data)   // boyut
}

// Map'deki herhangi bir anahtara ait veriyi gerçekleştir
func (s *SafeMap[K, V]) ForEach(f func(K, V)) {
	s.mu.RLock()               // mutex kilitlendi
	defer s.mu.RUnlock()       // mutex kilitlendi
	for k, v := range s.data { // herhangi bir anahtar ve veri okundu
		f(k, v) // fonksiyon cagrildi
	}
}
