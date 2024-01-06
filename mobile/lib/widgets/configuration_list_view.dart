import 'package:flutter/material.dart';
import 'package:tech_tinker/models/configuration.dart';
import 'package:tech_tinker/screens/configuration_screen.dart';
import 'package:tech_tinker/widgets/configuration_card.dart';

class ConfigurationListView extends StatefulWidget {
  final List<Configuration> configurations;

  const ConfigurationListView({
    super.key,
    required this.configurations,
  });

  @override
  State<ConfigurationListView> createState() => _ConfigurationListViewState();
}

class _ConfigurationListViewState extends State<ConfigurationListView> {
  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      shrinkWrap: true,
      itemBuilder: _configurationCardBuilder,
      itemCount: widget.configurations.length,
    );
  }

  Widget? _configurationCardBuilder(BuildContext context, int index) {
    return ConfigurationCard(
      number: index + 1,
      configuration: widget.configurations[index],
      onDelete: () {
        setState(() {
          widget.configurations.removeAt(index);
        });
      },
      onTap: () {
        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (context) => ConfigurationScreen(
              number: index + 1,
              configuration: widget.configurations[index],
            ),
          ),
        );
      },
    );
  }
}
