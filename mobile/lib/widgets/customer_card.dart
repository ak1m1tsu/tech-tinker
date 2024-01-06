import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/models/customer.dart';

class CustomerCard extends StatelessWidget {
  final Customer customer;

  const CustomerCard({
    super.key,
    required this.customer,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      color: CupertinoColors.extraLightBackgroundGray,
      child: ListTile(
        leading: const Icon(
          CupertinoIcons.person_fill,
          size: 24,
        ),
        title: Text(
          "${customer.firstName} ${customer.lastName}",
          style: const TextStyle(
            fontWeight: FontWeight.bold,
            fontSize: 18,
          ),
        ),
        onTap: () {},
      ),
    );
  }
}
