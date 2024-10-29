package dice

import (
    "net/http"
	"math/rand/v2"
	"encoding/json"
)

func DiceRoll (w http.ResponseWriter, r *http.Request) {
	result := rand.IntN(6) + 1;
	response := map[string]int{"message": result}
    json.NewEncoder(w).Encode(response)
}