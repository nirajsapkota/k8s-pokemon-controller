/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	pokemonv1 "github.com/nsapkota/k8s-pokemon-controller/api/v1"
)

// PokemonReconciler reconciles a Pokemon object
type PokemonReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=pokemon.nsapkota.dev,resources=pokemons,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pokemon.nsapkota.dev,resources=pokemons/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pokemon.nsapkota.dev,resources=pokemons/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Pokemon object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *PokemonReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// I don't want any context logging at the moment
	// _ = log.FromContext(ctx)

	var l = log.Log

	var pokemon pokemonv1.Pokemon
	if err := r.Get(ctx, req.NamespacedName, &pokemon); err != nil {
		l.Error(err, fmt.Sprintf("The pokemon '%s' does not exist in the namespace '%s'", pokemon.Spec.Name, pokemon.Namespace))
		return ctrl.Result{}, nil
	}

	l.Info(fmt.Sprintf("Pokedex *Beep boop*: %s is a %s type Pokemon. How fascinating!", pokemon.Spec.Name, strings.Join(pokemon.Spec.Type, ", ")))
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PokemonReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pokemonv1.Pokemon{}).
		Complete(r)
}
