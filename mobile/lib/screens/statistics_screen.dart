import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:syncfusion_flutter_charts/charts.dart';
import 'package:tech_tinker/models/statistics.dart';
import 'package:tech_tinker/widgets/budget_list_view.dart';
import 'package:tech_tinker/widgets/ruble_formatter.dart';
import 'package:tech_tinker/widgets/screen.dart';

class StatisticsScreen extends StatefulWidget {
  const StatisticsScreen({
    super.key,
  });

  @override
  State<StatisticsScreen> createState() => _StatisticsScreenState();
}

class _StatisticsScreenState extends State<StatisticsScreen> {
  late List<_PieData> _chartSource;
  late Statistics statistics;

  final List<Color> _chartSourceColors = [
    CupertinoColors.systemBlue,
    CupertinoColors.systemGreen,
    CupertinoColors.systemYellow,
    CupertinoColors.systemRed,
  ];

  @override
  void initState() {
    statistics = Statistics.fromRawJson("""{
      "budgets": [
        {
          "count": 6,
          "type": "Lower than 50K"
        },
        {
          "count": 2,
          "type": "Between 50K and 100K"
        },
        {
          "count": 2,
          "type": "Between 100K and 500K"
        },
        {
          "count": 1,
          "type": "Greater than 500K"
        }
      ],
      "from": "2024-01-06T15:29:01+08:00",
      "to": "2024-02-05T15:29:01+08:00",
      "total": 182000000
    }""");
    _chartSource = statistics.budgets
        .map(
          (e) => _PieData(
            e.type,
            e.count,
          ),
        )
        .toList();

    print("FROM ${statistics.from} | TO ${statistics.to}");
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Screen(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            IconButton(
              icon: const Icon(
                Icons.refresh,
                size: 32,
                color: CupertinoColors.darkBackgroundGray,
              ),
              onPressed: () {},
            ),
            Text(
              "${DateFormat("dd.MM.yyyy").format(statistics.from)} - ${DateFormat("dd.MM.yyyy").format(statistics.to)}",
              style: const TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 18,
              ),
            ),
            IconButton(
              icon: const Icon(
                Icons.date_range,
                size: 32,
                color: CupertinoColors.darkBackgroundGray,
              ),
              onPressed: () {},
            )
          ],
        ),
        const SizedBox(
          height: 10,
        ),
        SfCircularChart(
          title: ChartTitle(
            text: "Total: ${RubleFormatter.format(statistics.total / 100)}",
            textStyle: const TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 18,
            ),
          ),
          series: <PieSeries>[
            PieSeries<_PieData, String>(
              strokeColor: CupertinoColors.darkBackgroundGray,
              pointColorMapper: (_PieData _, int index) =>
                  _chartSourceColors[index],
              sortingOrder: SortingOrder.values.last,
              radius: '100%',
              dataSource: _chartSource,
              xValueMapper: (_PieData data, _) => data.xData,
              yValueMapper: (_PieData data, _) => data.yData,
            ),
          ],
        ),
        const SizedBox(
          height: 10,
        ),
        BudgetListView(
          budgets: statistics.budgets,
          colors: _chartSourceColors,
        ),
      ],
    );
  }
}

class _PieData {
  _PieData(this.xData, this.yData);

  final String xData;
  final int yData;
}
