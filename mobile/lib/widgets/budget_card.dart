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
        color: CupertinoColors.extraLightBackgroundGray,
        child: ListTile(
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
          leading: Icon(
            CupertinoIcons.device_desktop,
            size: 28,
            color: color,
          ),
          title: Text(
            budget.type,
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
              color: color,
            ),
          ),
          trailing: Text(
            "${budget.count}",
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
              color: color,
            ),
          ),
        ),
      ),
    );
  }
}
