import 'package:flutter/material.dart';
import '../data/action_api.dart';


class ActionScreen extends StatefulWidget {
  const ActionScreen({super.key});

  @override
  State<ActionScreen> createState() => _ActionScreenState();
}

class _ActionScreenState extends State<ActionScreen> with SingleTickerProviderStateMixin {
  // Tab controller for switching between investment and events
  late TabController _tabController;
  
  // Controllers for investment form
  final TextEditingController _fundIdController = TextEditingController();
  final TextEditingController _assetIdController = TextEditingController();
  final TextEditingController _amountController = TextEditingController();
  final TextEditingController _priceController = TextEditingController();
  
  // Fund ID value
  int _fundIdValue = 0;
  int _assetIdValue = 0;
  
  // Events data
  List<Event> events = [];
  List<AssetInfo>? assets = [];
  bool isLoading = false;
  final ActionApi actionApi = ActionApiProvider.getApi();

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 2, vsync: this);
    // Initialize the fund ID controller with the default value

    // Fetch events when screen loads
    fetchEvents();
    fetchAssets();
  }

  @override
  void dispose() {
    _tabController.dispose();
    _fundIdController.dispose();
    _assetIdController.dispose();
    _amountController.dispose();
    _priceController.dispose();
    super.dispose();
  }

  Future<void> fetchAssets() async {
  
   try {
      assets = await actionApi.getAssetInfos();
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Error fetching assets: ${e.toString()}')),
      );
    }
  }

  // Show the Fund ID selection dialog with only options 1, 2, 3
  void _showFundIdPicker() {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Select Fund ID'),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              ListTile(
                title: const Text('1'),
                selected: _fundIdValue == 1,
                onTap: () {
                  setState(() {
                    _fundIdValue = 1;
                    _fundIdController.text = '1';
                  });
                  Navigator.of(context).pop();
                },
              ),
              ListTile(
                title: const Text('2'),
                selected: _fundIdValue == 2,
                onTap: () {
                  setState(() {
                    _fundIdValue = 2;
                    _fundIdController.text = '2';
                  });
                  Navigator.of(context).pop();
                },
              ),
              ListTile(
                title: const Text('3'),
                selected: _fundIdValue == 3,
                onTap: () {
                  setState(() {
                    _fundIdValue = 3;
                    _fundIdController.text = '3';
                  });
                  Navigator.of(context).pop();
                },
              ),
            ],
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('CANCEL'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

    // Show the Fund ID selection dialog with only options 1, 2, 3
  void _showAssetIdPicker() {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Select Asset'),
          content: SizedBox(
            width: double.maxFinite, // Ensures width is flexible
          child: SingleChildScrollView(
            child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              ...assets!.map(
                (asset) => ListTile(
                 title: Text([asset.id, asset.name].join(". ")),
                 selected: _assetIdValue == asset.id,
                 onTap: (){
                  setState(() {
                    _assetIdValue = asset.id;
                    _assetIdController.text = _assetIdValue.toString();
                  });
                  Navigator.of(context).pop();
                 },
                ),
              ),
            ],
          ),
          ),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('CANCEL'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  Future<void> sendInvestmentAction() async {
    try {
      // Validate form data
      final fundId = int.tryParse(_fundIdController.text);
      final assetId = int.tryParse(_assetIdController.text);
      final amount = double.tryParse(_amountController.text);
      final price = double.tryParse(_priceController.text);
      
      if (fundId == null || assetId == null || amount == null || price == null) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Please enter valid values')),
        );
        return;
      }

      // For now, just simulate a successful response
      await actionApi.recordInvest(fundId, assetId, amount, price);
      
      // Clear form
      _fundIdController.clear();
      _assetIdController.clear();
      _amountController.clear();
      _priceController.clear();
      
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Investment action recorded successfully')),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Error: ${e.toString()}')),
      );
    }
  }

  Future<void> fetchEvents() async {
    setState(() {
      isLoading = true;
    });
    
    try {
      events = await actionApi.getEvents();
      setState(() {
        isLoading = false;
      });
    } catch (e) {
      setState(() {
        isLoading = false;
      });
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Error fetching events: ${e.toString()}')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Actions'),
        bottom: TabBar(
          controller: _tabController,
          tabs: const [
            Tab(text: 'Investment', icon: Icon(Icons.monetization_on)),
            Tab(text: 'Events', icon: Icon(Icons.event)),
          ],
        ),
      ),
      body: TabBarView(
        controller: _tabController,
        children: [
          // Investment recording form
          GestureDetector(
            onTap: () {
              FocusScope.of(context).unfocus(); // Hides keyboard when tapping outside
            },
            child: SingleChildScrollView(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  const Text(
                    'Record Investment Action',
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 16),
                  
                  // Fund ID - Now with limited options (1, 2, 3)
                  GestureDetector(
                    onTap: _showFundIdPicker,
                    child: AbsorbPointer(
                      child: TextField(
                        controller: _fundIdController,
                        decoration: const InputDecoration(
                          labelText: 'Fund ID',
                          border: OutlineInputBorder(),
                          suffixIcon: Icon(Icons.arrow_drop_down),
                          hintText: 'Click to select fund id',
                        ),
                        keyboardType: TextInputType.number,
                      ),
                    ),
                  ),
                  const SizedBox(height: 12),
                  GestureDetector(
                    onTap: _showAssetIdPicker,
                    child: AbsorbPointer(
                      child: TextField(
                      controller: _assetIdController,
                      decoration: const InputDecoration(
                        labelText: 'Asset ID',
                        border: OutlineInputBorder(),
                        suffixIcon: Icon(Icons.arrow_drop_down),
                          hintText: 'Click to select asset id',
                      ),
                      keyboardType: TextInputType.number,
                    ),
                    )
                  ),
                  
                  const SizedBox(height: 12),
                  TextField(
                    controller: _amountController,
                    decoration: const InputDecoration(
                      labelText: 'Amount',
                      border: OutlineInputBorder(),
                    ),
                    keyboardType: const TextInputType.numberWithOptions(decimal: true),
                  ),
                  const SizedBox(height: 12),
                  TextField(
                    controller: _priceController,
                    decoration: const InputDecoration(
                      labelText: 'Price',
                      border: OutlineInputBorder(),
                    ),
                    keyboardType: const TextInputType.numberWithOptions(decimal: true),
                  ),
                  const SizedBox(height: 24),
                  ElevatedButton(
                    onPressed: sendInvestmentAction,
                    style: ElevatedButton.styleFrom(
                      padding: const EdgeInsets.symmetric(vertical: 16),
                    ),
                    child: const Text('SEND', style: TextStyle(fontSize: 16)),
                  ),
                ],
              ),
            ),
          ),
          // Events list
          isLoading
              ? const Center(child: CircularProgressIndicator())
              : events.isEmpty
                  ? Center(
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          const Text('No events found'),
                          const SizedBox(height: 16),
                          ElevatedButton(
                            onPressed: fetchEvents,
                            child: const Text('Refresh'),
                          ),
                        ],
                      ),
                    )
                  : RefreshIndicator(
                      onRefresh: fetchEvents,
                      child: ListView.builder(
                        padding: const EdgeInsets.all(16),
                        itemCount: events.length,
                        itemBuilder: (context, index) {
                          final event = events[index];
                          return Card(
                            margin: const EdgeInsets.only(bottom: 16),
                            elevation: 2,
                            child: Padding(
                              padding: const EdgeInsets.all(16),
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Row(
                                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                    children: [
                                      Text(
                                        event.title,
                                        style: const TextStyle(
                                          fontSize: 18,
                                          fontWeight: FontWeight.bold,
                                        ),
                                      ),
                                      Container(
                                        padding: const EdgeInsets.symmetric(
                                          horizontal: 8,
                                          vertical: 4,
                                        ),
                                        decoration: BoxDecoration(
                                          color: _getStatusColor(event.status),
                                          borderRadius: BorderRadius.circular(12),
                                        ),
                                        child: Text(
                                          event.status,
                                          style: const TextStyle(
                                            color: Colors.white,
                                            fontSize: 12,
                                          ),
                                        ),
                                      ),
                                    ],
                                  ),
                                  const SizedBox(height: 8),
                                  Text(event.description),
                                  const SizedBox(height: 16),
                                  Row(
                                    mainAxisAlignment: MainAxisAlignment.end,
                                    children: [
                                      OutlinedButton.icon(
                                        onPressed: () async {
                                          final result = await actionApi.toggleEventStatus(event.id, event.status);
                                          if (result) {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                              const SnackBar(
                                                content: Text('Event switch successfully!'),
                                                backgroundColor: Colors.green,
                                              ),
                                            );
                                            await fetchEvents();
                                          } else {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                              const SnackBar(
                                                content: Text('Failed to run event.'),
                                                backgroundColor: Colors.red,
                                              ),
                                            );
                                          }
                                        },
                                        icon: const Icon(Icons.swap_horiz),
                                        label: const Text('SWITCH'),
                                      ),
                                      const SizedBox(width: 8),
                                      ElevatedButton.icon(
                                        onPressed: () async {
                                          final result = await actionApi.runEvent(event.id);
                                          if (result) {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                              const SnackBar(
                                                content: Text('Event ran successfully!'),
                                                backgroundColor: Colors.green,
                                              ),
                                            );
                                          } else {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                              const SnackBar(
                                                content: Text('Failed to run event.'),
                                                backgroundColor: Colors.red,
                                              ),
                                            );
                                          }
                                        },
                                        icon: const Icon(Icons.play_arrow),
                                        label: const Text('RUN NOW'),
                                      ),
                                    ],
                                  ),
                                ],
                              ),
                            ),
                          );
                        },
                      ),
                    ),
        ],
      ),
    );
  }

  Color _getStatusColor(String status) {
    switch (status.toLowerCase()) {
      case 'active':
        return Colors.green;
      case 'pending':
        return Colors.orange;
      // case 'completed':
      //   return Colors.blue;
      // case 'inactive':
      //   return Colors.grey;
      default:
        return Colors.blueGrey;
    }
  }
}

// Model class for Event
class Event {
  final int id;
  final String title;
  String status;
  String description;

  Event({
    required this.id,
    required this.title,
    required this.status,
    required this.description,
  });

  factory Event.fromJson(Map<String, dynamic> json) {
    return Event(
      id: json['id'],
      title: json['title'],
      status: json['status'],
      description: json['description'],
    );
  }
}

class AssetInfo {
  final int id;
  final String name;

  AssetInfo({
    required this.id,
    required this.name,
  });

  factory AssetInfo.fromJson(Map<String, dynamic> json) {
    return AssetInfo(
      id: json['id'],
      name: json['name'],
    );
  }
}