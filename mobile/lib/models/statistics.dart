import 'dart:convert';

class Statistics {
  final List<Budget> budgets;
  final DateTime from;
  final DateTime to;
  final int total;

  Statistics({
    required this.budgets,
    required this.from,
    required this.to,
    required this.total,
  });

  factory Statistics.fromRawJson(String str) =>
      Statistics.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Statistics.fromJson(Map<String, dynamic> json) => Statistics(
        budgets:
            List<Budget>.from(json["budgets"].map((x) => Budget.fromJson(x))),
        from: DateTime.parse(json["from"]),
        to: DateTime.parse(json["to"]),
        total: json["total"],
      );

  Map<String, dynamic> toJson() => {
        "budgets": List<dynamic>.from(budgets.map((x) => x.toJson())),
        "from": from.toIso8601String(),
        "to": to.toIso8601String(),
        "total": total,
      };
}

class Budget {
  final int count;
  final String type;

  Budget({
    required this.count,
    required this.type,
  });

  factory Budget.fromRawJson(String str) => Budget.fromJson(json.decode(str));

  String toRawJson() => json.encode(toJson());

  factory Budget.fromJson(Map<String, dynamic> json) => Budget(
        count: json["count"],
        type: json["type"],
      );

  Map<String, dynamic> toJson() => {
        "count": count,
        "type": type,
      };
}
