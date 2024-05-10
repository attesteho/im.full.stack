package controllers

import (
	"html/template"
	"net/http"
)

// Gestionnaire de la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "accueil.html")
}

// Gestionnaire de la page "services"
func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "services.html")
}

// Gestionnaire de la page "projets"
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}

// Gestionnaire de la page "contact"
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

// Fonction utilitaire pour rendre les modèles HTML
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Ajoute le préfixe "views/" au nom du modèle
	tmpl = "views/" + tmpl

	// Charge le modèle depuis le disque dur
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		// En cas d'erreur, renvoie une erreur interne au client
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécute le modèle et écrit sa sortie dans w
	err = t.Execute(w, nil)
	if err != nil {
		// En cas d'erreur, renvoie une erreur interne au client
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
