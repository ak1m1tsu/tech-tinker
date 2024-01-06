import 'package:flutter/cupertino.dart';
import 'package:tech_tinker/models/statistics.dart';
import 'package:tech_tinker/widgets/budget_card.dart';

class BudgetListView extends StatelessWidget {
  final List<Budget> budgets;
  final List<Color> colors;

  const BudgetListView({
    super.key,
    required this.budgets,
    required this.colors,
  });

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      shrinkWrap: true,
      itemBuilder: _budgetCardBuilder,
      itemCount: budgets.length,
    );
  }

  Widget? _budgetCardBuilder(BuildContext context, int index) {
    return BudgetCard(
      budget: budgets[index],
      color: colors[index],
    );
  }
}
