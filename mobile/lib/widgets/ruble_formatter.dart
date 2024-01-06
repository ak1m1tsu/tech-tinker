import 'package:money_formatter/money_formatter.dart';

class RubleFormatter {
  static String format(double amount) {
    return MoneyFormatter(
      amount: amount,
      settings: MoneyFormatterSettings(
        symbol: "₽",
      ),
    ).output.symbolOnRight;
  }
}
