:- dynamic usado/1.

entrada(guacamole, 200).
entrada(ensalada, 150).
entrada(consome, 300).
entrada(tostadas_caprese, 250).
carne(filete_de_cerdo, 400).
carne(pollo_al_horno, 280).
carne(carne_en_salsa, 320).
pescado(tilapia, 160).
pescado(salmon, 300).
pescado(trucha, 225).
postre(flan, 200).
postre(nueces_con_miel, 500).
postre(naranja_confitada, 450).
postre(flan_de_coco, 375).


usado(_) :- fail.

comida_completa(Entrada, PlatoPrincipal, Postre, Calorias) :-
    entrada(Entrada, CalEntrada),
    (carne(PlatoPrincipal, CalPlato) ; pescado(PlatoPrincipal, CalPlato)),
    postre(Postre, CalPostre),
    not(usado(Entrada)),
    not(usado(PlatoPrincipal)),
    not(usado(Postre)),
    Calorias is CalEntrada + CalPlato + CalPostre,
    asserta(usado(Entrada)),
    asserta(usado(PlatoPrincipal)),
    asserta(usado(Postre)).


combinaciones_posibles(MaxCalorias, Combinaciones) :-
    findall((Entrada, PlatoPrincipal, Postre, Calorias),
            (comida_completa(Entrada, PlatoPrincipal, Postre, Calorias),
            Calorias =< MaxCalorias),
            Combinaciones),
    retractall(usado(_)).

%combinaciones_posibles(1200, Combinaciones).
