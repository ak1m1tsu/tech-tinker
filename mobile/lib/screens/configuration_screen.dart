import 'package:flutter/material.dart';
import 'package:tech_tinker/models/configuration.dart';
import 'package:tech_tinker/widgets/app_bar.dart';

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
      appBar: techTinkerAppBar(
        context,
        "Configuration #$number",
      ),
    );
  }
}
