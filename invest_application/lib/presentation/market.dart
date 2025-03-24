import 'package:flutter/material.dart';
import '../data/market_api.dart';
import 'package:fl_chart/fl_chart.dart';

class MarketScreen extends StatefulWidget {
  const MarketScreen({super.key});

  @override
  State<MarketScreen> createState() => _MarketScreenState();
}

class _MarketScreenState extends State<MarketScreen> {
  final MarketApi _marketApi = MarketApiProvider.getApi();
  Map<String, dynamic>? _fearGreedData;
  Map<String, dynamic>? _nasdaqData;
  Map<String, dynamic>? _sp500Data;
  MarketStatus? _currentStatus = MarketStatus.BEAR;

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
void _showStatusSelectionDialog(BuildContext context) {
  MarketStatus? selectedStatus = _currentStatus;
  
  showDialog(
    context: context,
    builder: (BuildContext context) {
      return AlertDialog(
        title: const Text('Select Market Status'),
        content: StatefulBuilder(
          builder: (BuildContext context, StateSetter setState) {
            return SizedBox(
              width: double.maxFinite,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: MarketStatus.values.map((status) {
                  return RadioListTile<MarketStatus>(
                    title: Row(
                      children: [
                        Container(
                          width: 16,
                          height: 16,
                          decoration: BoxDecoration(
                            color: status.color,
                            shape: BoxShape.circle,
                          ),
                        ),
                        const SizedBox(width: 8),
                        Text(status.displayName),
                      ],
                    ),
                    value: status,
                    groupValue: selectedStatus,
                    onChanged: (MarketStatus? value) {
                      setState(() {
                        selectedStatus = value;
                      });
                    },
                  );
                }).toList(),
              ),
            );
          },
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(context).pop(),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () {
              if (selectedStatus != null) {
                _updateMarketStatus(selectedStatus!);
              }
              Navigator.of(context).pop();
            },
            child: const Text('OK'),
          ),
        ],
      );
    },
  );
}

Future<void> _updateMarketStatus(MarketStatus newStatus) async {
  try {
    // Show loading indicator (optional)
    // e.g. setState(() { _isLoading = true; });
    
    // await updateMarketStatus(newStatus); // todo
    
    // Update local state after successful API call
    setState(() {
      _currentStatus = newStatus;
    });
    
    // Show success message (optional)
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(content: Text('Market status updated successfully')),
    );
  } catch (e) {
    // Show error message
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text('Failed to update market status: ${e.toString()}')),
    );
  } finally {
    // Hide loading indicator (optional)
    // e.g. setState(() { _isLoading = false; });
  }
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
                    elevation: 4,
                    child: Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          const Text(
                            'Current Market Status',
                            style: TextStyle(
                              fontSize: 18,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          const SizedBox(height: 8), // Add spacing between the lines
                          if (_currentStatus != null)
                            Row(
                              children: [
                                Container(
                                  width: 16,
                                  height: 16,
                                  decoration: BoxDecoration(
                                    color: _currentStatus!.color,
                                    shape: BoxShape.circle,
                                  ),
                                ),
                                const SizedBox(width: 8),
                                GestureDetector(
                                  onTap: () => _showStatusSelectionDialog(context),
                                  child: Row(
                                    children: [
                                      Text(
                                        _currentStatus!.displayName,
                                        style: TextStyle(
                                          fontSize: 24,
                                          fontWeight: FontWeight.bold,
                                          color: _currentStatus!.color,
                                        ),
                                      ),
                                      const SizedBox(width: 4),
                                      const Icon(Icons.edit, size: 16),
                                    ],
                                  ),
                                ),
                              ],
                            ),
                        ],
                      ),
                    ),
                  ),
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


