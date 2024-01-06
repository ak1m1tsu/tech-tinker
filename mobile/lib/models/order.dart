import 'dart:convert';

import 'configuration.dart';
import 'customer.dart';

class Order {
  final String address;
  final String comment;
  final List<Configuration> configurations;
  final DateTime createdAt;
  final Customer customer;
  final String id;
  final int number;
  final int priceLimit;
  final String status;

  Order({
    required this.address,
    required this.comment,
    required this.configurations,
    required this.createdAt,
    required this.customer,
    required this.id,
    required this.number,
    required this.priceLimit,
    required this.status,
  });

  factory Order.fromRawJson(String str) => Order.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Order.fromJson(Map<String, dynamic> json) => Order(
        address: json["address"],
        comment: json["comment"],
        configurations: List<Configuration>.from(
            json["configurations"].map((x) => Configuration.fromJson(x))),
        createdAt: DateTime.parse(json["created_at"]),
        customer: Customer.fromJson(json["customer"]),
        id: json["id"],
        number: json["number"],
        priceLimit: json["price_limit"],
        status: json["status"],
      );

  Map<String, dynamic> toJson() => {
        "address": address,
        "comment": comment,
        "configurations":
            List<dynamic>.from(configurations.map((x) => x.toJson())),
        "created_at": createdAt.toIso8601String(),
        "customer": customer.toJson(),
        "id": id,
        "number": number,
        "price_limit": priceLimit,
        "status": status,
      };
}
