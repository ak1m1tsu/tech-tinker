import 'dart:convert';

import 'order.dart';

class Account {
  final DateTime createdAt;
  final String email;
  final String firstName;
  final String id;
  final String lastName;
  final List<Order> orders;
  final String role;

  Account({
    required this.createdAt,
    required this.email,
    required this.firstName,
    required this.id,
    required this.lastName,
    required this.orders,
    required this.role,
  });

  factory Account.fromRawJson(String str) => Account.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Account.fromJson(Map<String, dynamic> json) => Account(
        createdAt: DateTime.parse(json["created_at"]),
        email: json["email"],
        firstName: json["first_name"],
        id: json["id"],
        lastName: json["last_name"],
        orders: List<Order>.from(json["orders"].map((x) => Order.fromJson(x))),
        role: json["role"],
      );

  Map<String, dynamic> toJson() => {
        "created_at": createdAt.toIso8601String(),
        "email": email,
        "first_name": firstName,
        "id": id,
        "last_name": lastName,
        "orders": List<dynamic>.from(orders.map((x) => x.toJson())),
        "role": role,
      };
}
