package main
import "fmt"

var salary_levels = [4]int{10000, 40000, 50000, 100000};  
var salary_level_tax_percent = [4]int{5, 10, 15, 20};

func calculate_tax(salary int, salary_level_index int, tax int) int {
  if salary > 0 {
    if salary >= salary_levels[salary_level_index] {
      if salary_level_index == len(salary_levels) - 1 {
        tax += salary * salary_level_tax_percent[salary_level_index] / 100;
        return tax;
      }
      tax += salary_level_tax_percent[salary_level_index] * salary_levels[salary_level_index] / 100;
      salary = salary - salary_levels[salary_level_index];
    } else {
        tax += salary * salary_level_tax_percent[salary_level_index] / 100;
        salary = 0;
    }
    return calculate_tax(salary, salary_level_index + 1, tax);
  }
  return tax;
}

func main() {
  var salary int;
  fmt.Scan(&salary);
  fmt.Println(calculate_tax(salary, 0, 0));
}
