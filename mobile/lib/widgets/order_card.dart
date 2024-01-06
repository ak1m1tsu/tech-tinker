import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:money_formatter/money_formatter.dart';
import 'package:tech_tinker/models/order.dart';
import 'package:tech_tinker/screens/order_screen.dart';
import 'package:tech_tinker/widgets/order_status_badge.dart';

class OrderCard extends StatelessWidget {
  final Order order;

  const OrderCard({
    super.key,
    required this.order,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Card(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(10),
        ),
        color: CupertinoColors.extraLightBackgroundGray,
        child: ListTile(
          title: Text(
            "Order #${order.number}",
            style: const TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 20,
            ),
          ),
          subtitle: Text(
            MoneyFormatter(
              amount: order.priceLimit / 100,
              settings: MoneyFormatterSettings(
                symbol: "₽",
              ),
            ).output.symbolOnRight,
            style: const TextStyle(
              fontWeight: FontWeight.w600,
              fontSize: 16,
            ),
          ),
          trailing: OrderStatusBadge(status: order.status),
          onTap: () {
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) => OrderScreen(order: order),
              ),
            );
          },
        ),
      ),
    );
  }
}
