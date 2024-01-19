import 'dart:convert';

class LoginDetails {
  final String token;

  LoginDetails({
    required this.token,
  });

  factory LoginDetails.fromRawJson(String str) =>
      LoginDetails.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory LoginDetails.fromJson(Map<String, dynamic> json) => LoginDetails(
        token: json["token"],
      );

  Map<String, dynamic> toJson() => {
        "token": token,
      };
}
