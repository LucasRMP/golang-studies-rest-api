package personality

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdFromQuery(r *http.Request) int {
	variables := mux.Vars(r)
	idString := variables["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		return -1
	}

	return id
}

func FindAllController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(FindAll())
}

func FindByIdController(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromQuery(r)

	personality, err := FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func CreateController(w http.ResponseWriter, r *http.Request) {
	var personalityInput CreatePersonalityDTO
	json.NewDecoder(r.Body).Decode(&personalityInput)

	personality, err := Create(personalityInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func DeleteController(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromQuery(r)

	err := Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Personality deleted")
}

func UpdateController(w http.ResponseWriter, r *http.Request) {
	id := GetIdFromQuery(r)

	var personalityInput UpdatePersonalityDTO
	json.NewDecoder(r.Body).Decode(&personalityInput)

	personality, err := Update(id, personalityInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(personality)
}
