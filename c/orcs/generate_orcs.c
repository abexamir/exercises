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
    int initial_life1 = orc1->life;
    int initial_life2 = orc2->life;
    while (orc1->life > 0 && orc2->life > 0){
        orc2->life -= orc1->attack;
        orc1->life -= orc2->attack;
    }
    if (orc1->life <= 0 && orc2->life <= 0){
        // Both orcs are "dead", decide winner based on initial life
        return initial_life1 >= initial_life2 ? orc1 : orc2;
    }
    else if (orc1->life > 0){
        return orc1;
    }
    return orc2;
}

int main(){
    srand(time(NULL));
    Orc* army = generate_orcs(10);
    for (int i = 0; i < 10; i++){
        printf("Orc %d: attack: %d, life: %d\n", i, army[i].attack, army[i].life);
    }
    Orc* winner;
    int winner_index = -1;
    for (int i = 0; i < 10; i++){
        for (int j = 0; j < 10; j++){
            if (i == j){
                continue;
            }
            if (army[j].life > 0 && army[i].life > 0){
                winner = fight(&army[i], &army[j]);
                if (winner == &army[i]){
                    printf("Orc %d vs Orc %d: Orc %d wins, Orc %d dies\n", i, j, i, j);
                    if (winner->life > 0) {
                        winner_index = i;
                    }
                    else{
                        printf("Orc %d died after winning\n", i);
                    }
                } else if (winner == &army[j]){
                    printf("Orc %d vs Orc %d: Orc %d wins, Orc %d dies\n", i, j, j, i);
                    if (winner->life > 0) {
                        winner_index = j;
                    }
                    else{
                        printf("Orc %d died after winning\n", j);
                    }
                    break;
                }
                else {
                    exit(1);
                }
            }
        }
    }
    if (winner_index != -1) {
        printf("Orc %d is the winner\n", winner_index);
    }
    else{
        printf("No winner\n");
    }
    free(army);
    return 0;
 }





