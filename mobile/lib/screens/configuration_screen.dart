import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/models/configuration.dart';

class ConfigurationScreen extends StatelessWidget {
  final Configuration configuration;
  final int number;

  const ConfigurationScreen({
    super.key,
    required this.number,
    required this.configuration,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: CupertinoColors.darkBackgroundGray,
        leading: IconButton(
          icon: const Icon(
            Icons.arrow_back,
            color: CupertinoColors.extraLightBackgroundGray,
          ),
          onPressed: () {
            Navigator.pop(context);
          },
        ),
        title: Text(
          "Configuration #$number",
          style: const TextStyle(
            fontWeight: FontWeight.bold,
            fontSize: 24,
            color: CupertinoColors.extraLightBackgroundGray,
          ),
        ),
      ),
    );
  }
}
