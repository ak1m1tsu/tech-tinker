import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:tech_tinker/models/configuration.dart';
import 'package:tech_tinker/widgets/ruble_formatter.dart';

class ConfigurationCard extends StatefulWidget {
  final Configuration configuration;
  final int number;
  final VoidCallback onDelete;
  final GestureTapCallback? onTap;

  const ConfigurationCard({
    super.key,
    required this.number,
    required this.configuration,
    required this.onDelete,
    this.onTap,
  });

  @override
  State<ConfigurationCard> createState() => _ConfigurationCardState();
}

class _ConfigurationCardState extends State<ConfigurationCard> {
  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Card(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(10),
        ),
        color: CupertinoColors.extraLightBackgroundGray,
        child: ListTile(
          onTap: widget.onTap,
          leading: Text(
            "${widget.number}",
            style: const TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
            ),
          ),
          title: Text(
            RubleFormatter.format(widget.configuration.price / 100),
            style: const TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
            ),
          ),
          subtitle: Text(
            DateFormat().format(widget.configuration.createdAt),
            style: const TextStyle(
              fontWeight: FontWeight.w500,
              fontSize: 16,
            ),
          ),
          trailing: IconButton(
            icon: const Icon(
              CupertinoIcons.delete_simple,
              size: 28,
              color: CupertinoColors.systemRed,
            ),
            onPressed: widget.onDelete,
          ),
        ),
      ),
    );
  }
}
