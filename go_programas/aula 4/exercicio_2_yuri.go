package main

import "fmt"

type Arvore struct {
	info int
	esq *Arvore
	dir *Arvore
}

/* retorna a quantidade de nós que armazenam números pares */
func (a *Arvore) Pares () int {
	if(a == nil){
		return 0
	}
	cont := 0

	if(a.info % 2 == 0){
		cont++
	}
	cont = cont + a.esq.Pares()
	cont = cont + a.dir.Pares()
	return cont
}

/* retorna a quantidade de folhas de uma árvore binária */
func (a *Arvore) Folhas () int {
	if(a == nil){
		return 0
	}

	cont := 0

	if(a.esq == nil && a.dir == nil){
		cont++
	}

	cont = cont + a.esq.Folhas()
	cont = cont + a.dir.Folhas()
	return cont
}

/* compara se duas árvores binárias são iguais */
func (a *Arvore) Igual (b *Arvore) bool {
	if(a == nil || b == nil){
		return false
	}

	valor := false
	
	if(a.info == b.info){
		valor = true
	}
	
	if(a.esq != nil && b.esq != nil){ 
	valor = valor && a.esq.Igual(b.esq)
	}
	if(a.dir != nil && b.dir != nil){ 
	valor = valor && a.dir.Igual(b.dir)
	}

	return valor
}

func main () {
	a1 := &Arvore{4, nil, nil}
	a2 := &Arvore{2, nil, a1}
	a3 := &Arvore{5, nil, nil}
	a4 := &Arvore{6, nil, nil}
	a5 := &Arvore{3, a3, a4}
	a := &Arvore{1, a2, a5}
	b := &Arvore{7, a2, a5}
	fmt.Println(a.Pares() == 3)
	fmt.Println(a.Folhas() == 3)
	fmt.Println(a.Igual(a) == true)
	fmt.Println(a.Igual(b) == false)
}