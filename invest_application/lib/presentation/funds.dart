import 'package:flutter/material.dart';
import '../data/funds_api_mock.dart';
import 'package:fl_chart/fl_chart.dart';

class FundTableData {
  final String name;
  final String amount;
  final String amountDollar;
  final String profitRate;
  final String division;
  final String quantity;
  final String price;
  final String priceDollar;
  final bool isStable;

  FundTableData({
    required this.name,
    required this.amount,
    required this.amountDollar,
    required this.profitRate,
    required this.division,
    required this.quantity,
    required this.price,
    required this.priceDollar,
    required this.isStable,
  });
}

class Funds extends StatefulWidget {
  const Funds({super.key});

  @override
  State<Funds> createState() => _FundsState();
}

class _FundsState extends State<Funds> with SingleTickerProviderStateMixin {
  late List<FundData> fundsData;
  bool _sortAscending = false;
  int _sortColumnIndex = 1;
  late List<FundTableData> _sortedData;
  int? _selectedSection;
  bool _showDollar = false;

  List<PieChartSectionData> getSections() {
    return fundsData.asMap().entries.map((entry) {
      final index = entry.key;
      final data = entry.value;
      return PieChartSectionData(
        value: data.value,
        color: data.color,
        title: '${data.value}%',
        radius: _selectedSection == index ? 110 : 100,
        titleStyle: const TextStyle(
          fontSize: 16,
          fontWeight: FontWeight.bold,
          color: Colors.white,
        ),
        badgeWidget: _selectedSection == index 
            ? const Icon(Icons.check_circle, color: Colors.white)
            : null,
        badgePositionPercentageOffset: 0.98,
      );
    }).toList();
  }

  void _onPieChartSectionClicked(int index, Color color) {
    setState(() {
      if (_selectedSection == index) {
        _selectedSection = null;
        _sortedData = FundsApiMock.getFundsTableData();
      } else {
        _selectedSection = index;
        _sortedData = FundsApiMock.getFundsTableData()
            .where((data) => 
                (color == Colors.orange && data.isStable) ||
                (color == Colors.purple && !data.isStable))
            .toList();
      }
    });
  }

  @override
  void initState() {
    super.initState();
    fundsData = FundsApiMock.getFundsData();
    _sortedData = FundsApiMock.getFundsTableData();
    _sortColumnIndex = 1;
    _sortAscending = false;
    _sort(_sortColumnIndex, _sortAscending);
  }

  void _sort(int columnIndex, bool ascending) {
    setState(() {
      _sortColumnIndex = columnIndex;
      _sortAscending = ascending;

      _sortedData.sort((a, b) {
        if (columnIndex == 0 || columnIndex == 3) {
          var aValue = columnIndex == 0 ? a.name : a.division;
          var bValue = columnIndex == 0 ? b.name : b.division;
          return ascending
              ? aValue.compareTo(bValue)
              : bValue.compareTo(aValue);
        }

        var aValue = '';
        var bValue = '';
        
        switch (columnIndex) {
          case 1:
            aValue = _showDollar ? a.amountDollar : a.amount;
            bValue = _showDollar ? b.amountDollar : b.amount;
            break;
          case 2:
            aValue = a.profitRate;
            bValue = b.profitRate;
            break;
          case 4:
            aValue = a.quantity;
            bValue = b.quantity;
            break;
          case 5:
            aValue = _showDollar ? a.priceDollar : a.price;
            bValue = _showDollar ? b.priceDollar : b.price;
            break;
          default:
            return 0;
        }
        
        return ascending
            ? double.parse(aValue).compareTo(double.parse(bValue))
            : double.parse(bValue).compareTo(double.parse(aValue));
      });
    });
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: const Text('Funds Distribution'),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: () {
              setState(() {
                _sortedData = FundsApiMock.getFundsTableData();
                fundsData = FundsApiMock.getFundsData();
                _selectedSection = null;
              });
            },
          ),
          if (_selectedSection != null)
            IconButton(
              icon: const Icon(Icons.clear),
              onPressed: () {
                setState(() {
                  _selectedSection = null;
                  _sortedData = FundsApiMock.getFundsTableData();
                });
              },
            ),
        ],
      ),
      body: SingleChildScrollView(
        // padding: const EdgeInsets.only(bottom: 100), // 바텀 패딩 관련. 필요시 주석 해제
        child: Column(
          children: [
            SizedBox(
              height: 300,
              child: PieChart(
                PieChartData(
                  sections: getSections(),
                  sectionsSpace: 0,
                  centerSpaceRadius: 40,
                  pieTouchData: PieTouchData(
                    touchCallback: (FlTouchEvent event, pieTouchResponse) {
                      if (event is! FlTapUpEvent || 
                          pieTouchResponse == null || 
                          pieTouchResponse.touchedSection == null) return;
                      
                      final sectionIndex = pieTouchResponse.touchedSection!.touchedSectionIndex;
                      if (sectionIndex >= 0 && sectionIndex < fundsData.length) {
                        _onPieChartSectionClicked(sectionIndex, fundsData[sectionIndex].color);
                      }
                    },
                  ),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: SingleChildScrollView(
                scrollDirection: Axis.horizontal,
                child: DataTable(
                  sortAscending: _sortAscending,
                  sortColumnIndex: _sortColumnIndex,
                  columns: [
                    DataColumn(
                      label: const Text('Name'),
                      onSort: (columnIndex, ascending) => _sort(columnIndex, ascending),
                    ),
                    DataColumn(
                      label: const Text('Amount'),
                      onSort: (columnIndex, ascending) => _sort(columnIndex, ascending),
                    ),
                    DataColumn(
                      label: const Text('Profit Rate'),
                      onSort: (columnIndex, ascending) => _sort(columnIndex, ascending),
                    ),
                    DataColumn(
                      label: const Text('Division'),
                      onSort: (columnIndex, ascending) => _sort(columnIndex, ascending),
                    ),
                    DataColumn(
                      label: const Text('Quantity'),
                      onSort: (columnIndex, ascending) => _sort(columnIndex, ascending),
                    ),
                    DataColumn(
                      label: const Text('Price'),
                      onSort: (columnIndex, ascending) => _sort(columnIndex, ascending),
                    ),
                  ],
                  rows: _sortedData.map((data) => DataRow(
                    color: MaterialStateProperty.all(
                      data.isStable ? Colors.orange.withOpacity(0.2) : Colors.purple.withOpacity(0.2),
                    ),
                    cells: [
                      DataCell(Text(data.name)),
                      DataCell(
                        InkWell(
                          onTap: () {
                            if (double.parse(data.amountDollar) > 0) {
                              setState(() => _showDollar = !_showDollar);
                            }
                          },
                          child: Text(
                            double.parse(_showDollar ? data.amountDollar : data.amount) == 0 
                                ? '-' 
                                : (_showDollar ? '\$${data.amountDollar}' : '₩${data.amount}')
                          ),
                        ),
                      ),
                      DataCell(Text('${data.profitRate}%')),
                      DataCell(Text(data.division)),
                      DataCell(Text(data.quantity)),
                      DataCell(
                        InkWell(
                          onTap: () {
                            if (double.parse(data.priceDollar) > 0) {
                              setState(() => _showDollar = !_showDollar);
                            }
                          },
                          child: Text(
                            double.parse(_showDollar ? data.priceDollar : data.price) == 0
                                ? '-'
                                : (_showDollar ? '\$${data.priceDollar}' : '₩${data.price}')
                          ),
                        ),
                      ),
                    ],
                  )).toList(),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
