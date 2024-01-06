import 'dart:convert';

class Customer {
  final String email;
  final String firstName;
  final String id;
  final String lastName;
  final String phoneNumber;

  Customer({
    required this.email,
    required this.firstName,
    required this.id,
    required this.lastName,
    required this.phoneNumber,
  });

  factory Customer.fromRawJson(String str) =>
      Customer.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Customer.fromJson(Map<String, dynamic> json) => Customer(
        email: json["email"],
        firstName: json["first_name"],
        id: json["id"],
        lastName: json["last_name"],
        phoneNumber: json["phone_number"],
      );

  Map<String, dynamic> toJson() => {
        "email": email,
        "first_name": firstName,
        "id": id,
        "last_name": lastName,
        "phone_number": phoneNumber,
      };
}
