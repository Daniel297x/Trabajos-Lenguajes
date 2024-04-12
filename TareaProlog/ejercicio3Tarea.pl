contador([], _, 0).
contador(_, [], 0).
contador([C1|C1s], [C2|C2s], Total) :-
    (C1 \= C2 ->
        contador(C1s, C2s, A), Total is A + 1
    ;   contador(C1s, C2s, Total)
    ).



distanciaH(String1, String2, Total) :- string_chars(String1, Lcaracteres1),
    string_chars(String2, Lcaracteres2),
    contador(Lcaracteres1, Lcaracteres2, Total).

%distanciaH("romano","camino",X).
