#include <stdlib.h>
#include <stdio.h>

typedef struct {
    int index;
    int attack;
    int life;
} Orc;

Orc* generate_orcs(int amount){
    Orc* army = malloc(sizeof(Orc) * amount);
    for (int i = 0; i < amount; i++){
        army[i].index = i;
        army[i].attack = rand() % 10;
        army[i].life = rand() % 10;
    }
    return army;
}

Orc* fight(Orc* orc1, Orc* orc2){
    while (orc1->life > 0 && orc2->life > 0){
        orc2->life -= orc1->attack;
        orc1->life -= orc2->attack;
    }
    return orc1->life > 0 ? orc1 : orc2;
}

void print_orc(Orc* orc) {
    printf("Orc %d: attack: %d, life: %d\n", orc->index, orc->attack, orc->life);
}

int main(){
    srand(time(NULL));
    Orc* army = generate_orcs(10);
    for (int i = 0; i < 10; i++){
        print_orc(&army[i]);
    }
    Orc* winner;
    for (int i = 0; i < 10; i++){
        for (int j = i + 1; j < 10; j++){
            if (army[j].life > 0 && army[i].life > 0){
                winner = fight(&army[i], &army[j]);
                printf("Orc %d vs Orc %d: Orc %d wins, Orc %d dies\n", i, j, winner->index, winner == &army[i] ? j : i);
                if (winner->life <= 0) {
                    printf("Orc %d died after winning\n", winner->index);
                }
            }
        }
    }
    printf(winner && winner->life > 0 ? "Orc %d is the winner\n" : "No winner\n", winner->index);
    free(army);
    return 0;
}
