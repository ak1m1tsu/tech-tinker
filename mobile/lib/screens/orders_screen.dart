import 'package:flutter/material.dart';
import 'package:tech_tinker/models/configuration.dart';
import 'package:tech_tinker/models/customer.dart';
import 'package:tech_tinker/models/order.dart';
import 'package:tech_tinker/widgets/order_list_view.dart';
import 'package:tech_tinker/widgets/screen.dart';

class OrdersScreen extends StatefulWidget {
  const OrdersScreen({super.key});

  @override
  State<OrdersScreen> createState() => _OrdersScreenState();
}

class _OrdersScreenState extends State<OrdersScreen> {
  @override
  Widget build(BuildContext context) {
    return Screen(
      children: [
        OrderListView(
          orders: [
            Order(
              address: "Irkutsk, Donskaya Street, 4, 21",
              comment: "Please, make config without RGB lightning.",
              configurations: [
                Configuration(
                  createdAt: DateTime.now(),
                  id: "id",
                  price: 15000000,
                )
              ],
              createdAt: DateTime.now(),
              customer: Customer(
                id: 'id',
                email: 'ivan.ivanov@gmail.com',
                phoneNumber: '88005553535',
                firstName: 'Ivan',
                lastName: 'Ivanov',
              ),
              id: "id",
              number: 1,
              priceLimit: 15000000,
              status: "In Process",
            )
          ],
        ),
      ],
    );
  }
}
