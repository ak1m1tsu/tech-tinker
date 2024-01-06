import 'package:flutter/material.dart';
import 'package:tech_tinker/models/order.dart';
import 'package:tech_tinker/widgets/order_card.dart';

class OrderListView extends StatelessWidget {
  final List<Order> orders;

  const OrderListView({
    super.key,
    required this.orders,
  });

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      shrinkWrap: true,
      itemBuilder: _orderCardBuilder,
      itemCount: orders.length,
    );
  }

  Widget? _orderCardBuilder(BuildContext context, int index) {
    return OrderCard(
      order: orders[index],
    );
  }
}
