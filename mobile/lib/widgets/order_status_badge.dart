import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class OrderStatusBadge extends StatelessWidget {
  final String status;

  const OrderStatusBadge({
    super.key,
    required this.status,
  });

  @override
  Widget build(BuildContext context) {
    return Badge(
      backgroundColor: _colorFromStatus(status),
      largeSize: 20,
      padding: const EdgeInsets.symmetric(horizontal: 10),
      label: Text(
        status,
        style: const TextStyle(
          fontWeight: FontWeight.bold,
          fontSize: 14,
        ),
      ),
    );
  }

  Color _colorFromStatus(String status) {
    switch (status) {
      case "In Process":
        return CupertinoColors.systemBlue;
      case "Completed":
        return CupertinoColors.systemRed;
      default:
        return CupertinoColors.darkBackgroundGray;
    }
  }
}
