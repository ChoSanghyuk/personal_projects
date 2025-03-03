import 'package:flutter/material.dart';
import '../data/hist_api.dart';
import '../data/hist_api_mock.dart';

class InvestmentRecord {
  final DateTime date;
  final String name;
  final double price;
  final double amount;
  final String action; // "buy" or "sell"

  InvestmentRecord({
    required this.date,
    required this.name,
    required this.price,
    required this.amount,
    required this.action,
  });
}

class HistScreen extends StatefulWidget {
  const HistScreen({super.key});

  @override
  State<HistScreen> createState() => _HistScreenState();
}

class _HistScreenState extends State<HistScreen> with SingleTickerProviderStateMixin {
  final HistoryApi historyApi = HistoryApiHttp();
  late TabController _tabController;
  late DateTimeRange _selectedDateRange;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 3, vsync: this);
    // Initialize with last 30 days as default
    _selectedDateRange = DateTimeRange(
      start: DateTime.now().subtract(const Duration(days: 30)),
      end: DateTime.now(),
    );
  }

  Future<void> _showDateRangePicker() async {
    final DateTimeRange? picked = await showDateRangePicker(
      context: context,
      firstDate: DateTime(2020),
      lastDate: DateTime.now(),
      initialDateRange: _selectedDateRange,
    );
    if (picked != null) {
      setState(() {
        _selectedDateRange = picked;
      });
    }
  }

  @override
  void dispose() {
    _tabController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: const Text('History'),
        actions: [
          IconButton(
            icon: const Icon(Icons.date_range),
            onPressed: _showDateRangePicker,
            tooltip: 'Select Date Range',
          ),
        ],
        bottom: TabBar(
          controller: _tabController,
          tabs: const [
            Tab(text: 'Fund 1'),
            Tab(text: 'Fund 2'),
            Tab(text: 'Fund 3'),
          ],
        ),
      ),
      body: TabBarView(
        controller: _tabController,
        children: [
          Padding(
            padding: const EdgeInsets.only(bottom: 10), // 바텀 패딩 관련. 필요시 100
            child: _buildHistoryList(1),
          ),
          Padding(
            padding: const EdgeInsets.only(bottom: 10), // 바텀 패딩 관련. 필요시 100
            child: _buildHistoryList(2),
          ),
          Padding(
            padding: const EdgeInsets.only(bottom: 10), // 바텀 패딩 관련. 필요시 100
            child: _buildHistoryList(3),
          ),
        ],
      ),
    );
  }

  Widget _buildHistoryList(int fundId) {
    return FutureBuilder<List<InvestmentRecord>>(
      future: historyApi.getInvestmentHistory(fundId, _selectedDateRange),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Center(child: CircularProgressIndicator());
        }
        
        if (snapshot.hasError) {
          return Center(child: Text('Error: ${snapshot.error}'));
        }

        final history = snapshot.data ?? [];
        
        // Sort history by date (most recent first)
        history.sort((a, b) => b.date.compareTo(a.date));
        
        return ListView.builder(
          itemCount: history.length,
          itemBuilder: (context, index) {
            final record = history[index];
            return Card(
              margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
              child: ListTile(
                title: Text(record.name),
                subtitle: Text(
                  '${record.action.toUpperCase()} - ${record.amount} at \$${record.price}',
                ),
                trailing: Text(
                  '${record.date.day}/${record.date.month}/${record.date.year}',
                  style: Theme.of(context).textTheme.bodyMedium,
                ),
                leading: Icon(
                  record.action == 'BUY' ? Icons.arrow_downward : Icons.arrow_upward,
                  color: record.action == 'BUY' ? Colors.green : Colors.red,
                ),
              ),
            );
          },
        );
      },
    );
  }
}
