import 'package:flutter/material.dart';
import '../data/market_api.dart';
import '../data/market_api_mock.dart';
import 'package:fl_chart/fl_chart.dart';

class MarketScreen extends StatefulWidget {
  const MarketScreen({super.key});

  @override
  State<MarketScreen> createState() => _MarketScreenState();
}

class _MarketScreenState extends State<MarketScreen> {
  final MarketApi _marketApi = MarketApiHttp();
  Map<String, dynamic>? _fearGreedData;
  Map<String, dynamic>? _nasdaqData;
  Map<String, dynamic>? _sp500Data;

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  Future<void> _loadData() async {
    final loaded = await _marketApi.getIndexs();
    
    setState(() {
      _fearGreedData = loaded['fearGreed'];
      _nasdaqData = loaded['nasdaq'];
      _sp500Data = loaded['sp500'];
    });
  }

  // Widget _buildTrendGraph(List<dynamic> data, Color color) {
  //   final spots = List.generate(
  //     data.length,
  //     (i) => FlSpot(i.toDouble()+1, data[i].toDouble()),
  //   );

  //   print(spots);
  //   return LineChart(
  //     LineChartData(
  //       gridData: const FlGridData(show: false),
  //       titlesData: const FlTitlesData(show: false),
  //       borderData: FlBorderData(show: false),
  //       minX: spots.first.x,
  //     minY: spots.map((e) => e.y).reduce((a, b) => a < b ? a : b), // Find min y-value dynamically
  //     maxX: spots.last.x,
  //     maxY: spots.map((e) => e.y).reduce((a, b) => a > b ? a : b), // Find max y-value dynamically

  //       lineBarsData: [
  //         LineChartBarData(
  //           spots: spots,
  //           isCurved: true,
  //           color: color,
  //           dotData: const FlDotData(show: false),
  //           belowBarData: BarAreaData(
  //             show: true,
  //             color: color.withOpacity(0.1),
  //           ),
  //         ),
  //       ],
  //     ),
  //   );
  // }

  Widget _buildTrendGraph(List<dynamic> data, Color color) {
  final spots = List.generate(
    data.length,
    (i) => FlSpot(i.toDouble(), data[i].toDouble()), // Use actual data index for X
  );

  double minYValue = spots.map((e) => e.y).reduce((a, b) => a < b ? a : b); // Get lowest Y value
  double maxYValue = spots.map((e) => e.y).reduce((a, b) => a > b ? a : b); // Get highest Y value

  return LineChart(
    LineChartData(
      gridData: const FlGridData(show: false),
      titlesData: const FlTitlesData(show: false),
      borderData: FlBorderData(show: false),

      // ✅ Set dynamic minY with padding (so it doesn't stick to the bottom)
      minX: spots.first.x,
      minY: minYValue - (maxYValue - minYValue) * 0.1,  // Add 10% padding below
      maxX: spots.last.x,
      maxY: maxYValue + (maxYValue - minYValue) * 0.1,  // Add 10% padding above

      lineBarsData: [
        LineChartBarData(
          spots: spots,
          isCurved: true,
          color: color,
          dotData: const FlDotData(show: false),
          belowBarData: BarAreaData(
            show: true,
            color: color.withOpacity(0.1),
          ),
        ),
      ],
    ),
  );
}


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: const Text('Market'),
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 16.0), // 바텀 패딩 관련. 필요시 바텀 값 116
          child: RefreshIndicator(
            onRefresh: _loadData,
            child: SingleChildScrollView(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Card(
                    child: Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: _fearGreedData == null
                          ? const Center(child: CircularProgressIndicator())
                          : Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                const Text(
                                  'Fear & Greed Index',
                                  style: TextStyle(
                                    fontSize: 18,
                                    fontWeight: FontWeight.bold,
                                  ),
                                ),
                                const SizedBox(height: 8),
                                Row(
                                  children: [
                                    Text(
                                      '${_fearGreedData!['value']}',
                                      style: const TextStyle(
                                        fontSize: 24,
                                        fontWeight: FontWeight.bold,
                                        color: Colors.green,
                                      ),
                                    ),
                                    const SizedBox(width: 8),
                                    Container(
                                      padding: const EdgeInsets.symmetric(
                                        horizontal: 8,
                                        vertical: 4,
                                      ),
                                      decoration: BoxDecoration(
                                        color: Colors.green.withOpacity(0.1),
                                        borderRadius: BorderRadius.circular(4),
                                      ),
                                      child: Text(_fearGreedData!['status']),
                                    ),
                                  ],
                                ),
                                const SizedBox(height: 16),
                                Container(
                                  height: 100,
                                  width: double.infinity,
                                  decoration: BoxDecoration(
                                    border: Border.all(color: Colors.grey.shade300),
                                    borderRadius: BorderRadius.circular(8),
                                  ),
                                  child: _fearGreedData == null
                                      ? const Center(child: CircularProgressIndicator())
                                      : _buildTrendGraph(
                                          _fearGreedData!['weeklyData'],
                                          Colors.orange,
                                        ),
                                ),
                              ],
                            ),
                    ),
                  ),
                  const SizedBox(height: 16),
                  Card(
                    child: Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: _nasdaqData == null
                          ? const Center(child: CircularProgressIndicator())
                          : Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                const Text(
                                  'NASDAQ',
                                  style: TextStyle(
                                    fontSize: 18,
                                    fontWeight: FontWeight.bold,
                                  ),
                                ),
                                const SizedBox(height: 8),
                                Row(
                                  children: [
                                    Text(
                                      _nasdaqData!['value'].toString(),
                                      style: const TextStyle(
                                        fontSize: 24,
                                        fontWeight: FontWeight.bold,
                                      ),
                                    ),
                                    const SizedBox(width: 8),
                                    Container(
                                      padding: const EdgeInsets.symmetric(
                                        horizontal: 8,
                                        vertical: 4,
                                      ),
                                      decoration: BoxDecoration(
                                        color: Colors.green.withOpacity(0.1),
                                        borderRadius: BorderRadius.circular(4),
                                      ),
                                      child: Text(
                                        '${_nasdaqData!['change'].toStringAsFixed(2)}%',
                                        style: const TextStyle(color: Colors.green),
                                      ),
                                    ),
                                  ],
                                ),
                                const SizedBox(height: 16),
                                Container(
                                  height: 100,
                                  width: double.infinity,
                                  decoration: BoxDecoration(
                                    border: Border.all(color: Colors.grey.shade300),
                                    borderRadius: BorderRadius.circular(8),
                                  ),
                                  child: _nasdaqData == null
                                      ? const Center(child: CircularProgressIndicator())
                                      : _buildTrendGraph(
                                          _nasdaqData!['weeklyData'],
                                          Colors.blue,
                                        ),
                                ),
                              ],
                            ),
                    ),
                  ),
                  const SizedBox(height: 16),
                  Card(
                    child: Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: _sp500Data == null
                          ? const Center(child: CircularProgressIndicator())
                          : Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                const Text(
                                  'S&P 500',
                                  style: TextStyle(
                                    fontSize: 18,
                                    fontWeight: FontWeight.bold,
                                  ),
                                ),
                                const SizedBox(height: 8),
                                Row(
                                  children: [
                                    Text(
                                      _sp500Data!['value'].toString(),
                                      style: const TextStyle(
                                        fontSize: 24,
                                        fontWeight: FontWeight.bold,
                                      ),
                                    ),
                                    const SizedBox(width: 8),
                                    Container(
                                      padding: const EdgeInsets.symmetric(
                                        horizontal: 8,
                                        vertical: 4,
                                      ),
                                      decoration: BoxDecoration(
                                        color: Colors.green.withOpacity(0.1),
                                        borderRadius: BorderRadius.circular(4),
                                      ),
                                      child: Text(
                                        '${_sp500Data!['change'].toStringAsFixed(2)}%',
                                        style: const TextStyle(color: Colors.green),
                                      ),
                                    ),
                                  ],
                                ),
                                const SizedBox(height: 16),
                                Container(
                                  height: 100,
                                  width: double.infinity,
                                  decoration: BoxDecoration(
                                    border: Border.all(color: Colors.grey.shade300),
                                    borderRadius: BorderRadius.circular(8),
                                  ),
                                  child: _sp500Data == null
                                      ? const Center(child: CircularProgressIndicator())
                                      : _buildTrendGraph(
                                          _sp500Data!['weeklyData'],
                                          Colors.blue,
                                        ),
                                ),
                              ],
                            ),
                    ),
                  ),
                  // Card(
                  //   child: _sp500Data == null
                  //       ? const Center(child: CircularProgressIndicator())
                  //       : ListTile(
                  //           leading: const Icon(Icons.show_chart),
                  //           title: const Text('S&P 500'),
                  //           subtitle: Text(_sp500Data!['value'].toString()),
                  //           trailing: Container(
                  //             padding: const EdgeInsets.all(8),
                  //             decoration: BoxDecoration(
                  //               color: Colors.green.withOpacity(0.1),
                  //               borderRadius: BorderRadius.circular(8),
                  //             ),
                  //             child: Text(
                  //               '+${_sp500Data!['change']}%',
                  //               style: const TextStyle(color: Colors.green),
                  //             ),
                  //           ),
                  //         ),
                  // ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
