import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class SettingsCard extends StatelessWidget {
  SettingsCard({
    super.key,
    required this.text,
    required this.icon,
    required this.onTap,
    this.textColor = CupertinoColors.black,
    this.backgroundColor = CupertinoColors.white,
  });

  final String text;
  final IconData icon;
  final GestureTapCallback? onTap;

  Color backgroundColor;
  Color textColor;

  @override
  Widget build(BuildContext context) {
    return Card(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
      ),
      color: backgroundColor,
      child: ListTile(
        onTap: onTap,
        title: Text(
          text,
          style: TextStyle(
            fontWeight: FontWeight.w700,
            fontSize: 18,
            color: textColor,
          ),
        ),
        leading: Icon(
          icon,
          size: 30,
          color: textColor,
        ),
      ),
    );
  }
}
