import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/models/statistics.dart';

class BudgetCard extends StatelessWidget {
  final Budget budget;
  final Color color;

  const BudgetCard({
    super.key,
    required this.budget,
    required this.color,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Card(
        color: color,
        child: ListTile(
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
          leading: const Icon(
            CupertinoIcons.device_desktop,
            size: 28,
            color: CupertinoColors.darkBackgroundGray,
          ),
          title: Text(
            budget.type,
            style: const TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
              color: CupertinoColors.darkBackgroundGray,
            ),
          ),
          trailing: Text(
            "${budget.count}",
            style: const TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
              color: CupertinoColors.darkBackgroundGray,
            ),
          ),
        ),
      ),
    );
  }
}
