
concatenar([], L, L).
concatenar([X|L1], L2, [X|L3]) :- concatenar(L1, L2, L3).

aplanar([], Acumulador, Acumulador).
aplanar([X|Resto], Acumulador, Resultado) :-
    (   is_list(X) ->
        aplanar(X, [], ListaX),
        concatenar(ListaX, Acumulador, AcumuladorActualizado),
        aplanar(Resto, AcumuladorActualizado, Resultado);
        aplanar(Resto, [X|Acumulador], Resultado)
    ).

aplanar(Lista, Resultado) :- aplanar(Lista, [], Resultado).


%aplanar([1,2,[3,[4,5],[6,7]]],X).
%da el resultado alrevez, pero lo hace jajaj
