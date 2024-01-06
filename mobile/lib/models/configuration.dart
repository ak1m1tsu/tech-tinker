import 'dart:convert';

class Configuration {
  final DateTime createdAt;
  final String id;
  final int price;

  Configuration({
    required this.createdAt,
    required this.id,
    required this.price,
  });

  factory Configuration.fromRawJson(String str) =>
      Configuration.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Configuration.fromJson(Map<String, dynamic> json) => Configuration(
        createdAt: DateTime.parse(json["created_at"]),
        id: json["id"],
        price: json["price"],
      );

  Map<String, dynamic> toJson() => {
        "created_at": createdAt.toIso8601String(),
        "id": id,
        "price": price,
      };
}
