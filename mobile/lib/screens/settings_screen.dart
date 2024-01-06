import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/api/auth.dart';
import 'package:tech_tinker/widgets/screen.dart';
import 'package:tech_tinker/widgets/settings_card.dart';

class SettingsScreen extends StatelessWidget {
  const SettingsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Screen(
      children: [
        const Text(
          "General",
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.bold,
          ),
        ),
        const SizedBox(height: 10),
        SettingsCard(
          text: "Logout",
          icon: Icons.exit_to_app,
          onTap: () {
            AuthCache.logout();
            Navigator.pushNamedAndRemoveUntil(
              context,
              "/login",
              (route) => false,
            );
          },
          textColor: CupertinoColors.systemRed,
        ),
      ],
    );
  }
}
