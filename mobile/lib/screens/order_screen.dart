import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:money_formatter/money_formatter.dart';
import 'package:tech_tinker/models/order.dart';
import 'package:tech_tinker/widgets/configuration_list_view.dart';
import 'package:tech_tinker/widgets/customer_card.dart';

class OrderScreen extends StatefulWidget {
  final Order order;

  const OrderScreen({
    super.key,
    required this.order,
  });

  @override
  State<OrderScreen> createState() => _OrderScreenState();
}

class _OrderScreenState extends State<OrderScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: CupertinoColors.darkBackgroundGray,
        leading: IconButton(
          icon: const Icon(
            Icons.arrow_back,
            color: CupertinoColors.extraLightBackgroundGray,
          ),
          onPressed: () {
            Navigator.pop(context);
          },
        ),
        title: Text(
          "Order #${widget.order.number}",
          style: const TextStyle(
            fontWeight: FontWeight.bold,
            fontSize: 24,
            color: CupertinoColors.extraLightBackgroundGray,
          ),
        ),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(20),
        child: Column(
          children: [
            CustomerCard(customer: widget.order.customer),
            Card(
              color: CupertinoColors.extraLightBackgroundGray,
              child: Column(
                children: [
                  ListTile(
                    leading: const Icon(
                      Icons.comment,
                      size: 28,
                    ),
                    title: Text(
                      widget.order.comment,
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                      ),
                    ),
                  ),
                  ListTile(
                    leading: const Icon(
                      CupertinoIcons.location_solid,
                      size: 28,
                    ),
                    title: Text(
                      widget.order.address,
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Card(
              color: CupertinoColors.extraLightBackgroundGray,
              child: Column(
                children: [
                  ListTile(
                    leading: const Icon(
                      CupertinoIcons.money_rubl,
                      size: 28,
                    ),
                    title: Text(
                      MoneyFormatter(
                        amount: widget.order.priceLimit / 100,
                      ).output.nonSymbol,
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                      ),
                    ),
                  ),
                  ListTile(
                    leading: const Icon(
                      CupertinoIcons.star,
                      size: 28,
                    ),
                    title: Text(
                      widget.order.status,
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                      ),
                    ),
                  ),
                  ListTile(
                    leading: const Icon(
                      CupertinoIcons.time,
                      size: 28,
                    ),
                    title: Text(
                      DateFormat().format(widget.order.createdAt),
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                      ),
                    ),
                  ),
                ],
              ),
            ),
            const Padding(
              padding: EdgeInsets.symmetric(
                vertical: 10,
              ),
              child: Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Configurations",
                  style: TextStyle(
                    fontWeight: FontWeight.w500,
                    fontSize: 16,
                  ),
                ),
              ),
            ),
            ConfigurationListView(
              configurations: widget.order.configurations,
            )
          ],
        ),
      ),
    );
  }
}
